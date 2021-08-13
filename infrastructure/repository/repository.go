package repository

import (
	"context"
	"fmt"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	P     *db.Postgres
	Kafka *external.Kafka
}

func NewRepository(postgres *db.Postgres, kafka *external.Kafka) *Repository {
	return &Repository{
		P:     postgres,
		Kafka: kafka,
	}
}

func (r *Repository) CreateCompany(ctx context.Context, company *entity.Company) error {
	err := r.P.Db.Create(company).Error
	return err
}

func (r *Repository) FindCompany(ctx context.Context, id string) (*entity.Company, error) {
	var company entity.Company
	r.P.Db.First(&company, "id = ?", id)

	if company.ID == "" {
		return nil, fmt.Errorf("no company found")
	}

	return &company, nil
}

func (r *Repository) SearchCompanies(ctx context.Context, filter *entity.Filter) (*string, []*entity.Company, error) {
	var companies []*entity.Company

	q := r.P.Db.Order("token desc").Limit(filter.PageSize)

	if filter.CorporateName != "" {
		q = q.Where("corporate_name = ?", filter.CorporateName)
	}
	if filter.TradeName != "" {
		q = q.Where("trade_name = ?", filter.TradeName)
	}
	if filter.Cnpj != "" {
		q = q.Where("cnpj = ?", filter.Cnpj)
	}
	if filter.PageToken != "" {
		q = q.Where("token < ?", filter.PageToken)
	}

	err := q.Find(&companies).Error
	if err != nil {
		return nil, nil, err
	}

	length := len(companies)
	var nextPageToken string
	if length == filter.PageSize {
		nextPageToken = *companies[length-1].Token
	}

	return &nextPageToken, companies, nil
}

func (r *Repository) SaveCompany(ctx context.Context, company *entity.Company) error {
	err := r.P.Db.Save(company).Error
	return err
}

func (r *Repository) PublishEvent(ctx context.Context, msg, topic, key string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
		Key:            []byte(key),
	}
	err := r.Kafka.Producer.Produce(message, r.Kafka.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateEmployee(ctx context.Context, company *entity.Employee) error {
	err := r.P.Db.Create(company).Error
	return err
}

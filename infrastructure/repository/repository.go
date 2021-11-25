package repository

import (
	"context"
	"fmt"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/domain/entity/filter"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	P *db.Postgres
	K *external.KafkaProducer
}

func NewRepository(postgres *db.Postgres, kafkaProducer *external.KafkaProducer) *Repository {
	return &Repository{
		P: postgres,
		K: kafkaProducer,
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

func (r *Repository) SearchCompanies(ctx context.Context, companyFilter *filter.CompanyFilter) (*string, []*entity.Company, error) {
	var companies []*entity.Company

	q := r.P.Db.Order("token desc").Limit(companyFilter.PageSize)

	if companyFilter.CorporateName != "" {
		q = q.Where("corporate_name = ?", companyFilter.CorporateName)
	}
	if companyFilter.TradeName != "" {
		q = q.Where("trade_name = ?", companyFilter.TradeName)
	}
	if companyFilter.Cnpj != "" {
		q = q.Where("cnpj = ?", companyFilter.Cnpj)
	}
	if companyFilter.PageToken != "" {
		q = q.Where("token < ?", companyFilter.PageToken)
	}

	err := q.Find(&companies).Error
	if err != nil {
		return nil, nil, err
	}

	length := len(companies)
	var nextPageToken string
	if length == companyFilter.PageSize {
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
	err := r.K.Producer.Produce(message, r.K.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateEmployee(ctx context.Context, company *entity.Employee) error {
	err := r.P.Db.Create(company).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	var employee entity.Employee
	r.P.Db.First(&employee, "id = ?", id)

	if employee.ID == "" {
		return nil, fmt.Errorf("no employee found")
	}

	return &employee, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.P.Db.Save(employee).Error
	return err
}

func (r *Repository) AddEmployeeToCompany(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error {
	err := r.P.Db.Create(companyEmployee).Error
	return err
}

func (r *Repository) CreateWorkScale(ctx context.Context, workScale *entity.WorkScale) error {
	err := r.P.Db.Create(workScale).Error
	return err
}

func (r *Repository) FindWorkScale(ctx context.Context, companyID, workScaleID string) (*entity.WorkScale, error) {
	var workScale entity.WorkScale
	r.P.Db.Preload("Clocks").First(&workScale, "id = ? AND company_id = ?", workScaleID, companyID)

	if workScale.ID == "" {
		return nil, fmt.Errorf("no work scale found")
	}

	return &workScale, nil
}

func (r *Repository) SaveWorkScale(ctx context.Context, workScale *entity.WorkScale) error {
	err := r.P.Db.Save(workScale).Error
	return err
}

func (r *Repository) SearchWorkScales(ctx context.Context, workScaleFilter *filter.WorkScaleFilter) ([]*entity.WorkScale, error) {
	var workScales []*entity.WorkScale

	q := r.P.Db.Preload("Clocks").Where("company_id = ?", workScaleFilter.CompanyID)

	if workScaleFilter.Name != "" {
		q = q.Where("name = ?", workScaleFilter.Name)
	}

	err := q.Find(&workScales).Error
	if err != nil {
		return nil, err
	}

	return workScales, nil
}

func (r *Repository) CreateClock(ctx context.Context, clock *entity.Clock) error {
	err := r.P.Db.Create(clock).Error
	return err
}

func (r *Repository) FindClock(ctx context.Context, workScaleID, clockID string) (*entity.Clock, error) {
	var clock entity.Clock
	r.P.Db.First(&clock, "id = ? AND work_scale_id = ?", clockID, workScaleID)

	if clock.ID == "" {
		return nil, fmt.Errorf("no clock found")
	}

	return &clock, nil
}

func (r *Repository) DeleteClock(ctx context.Context, workScaleID, clockID string) error {
	err := r.P.Db.Where("id = ? AND work_scale_id = ?", clockID, workScaleID).Delete(entity.Clock{}).Error
	return err
}

func (r *Repository) SaveClock(ctx context.Context, clock *entity.Clock) error {
	err := r.P.Db.Save(clock).Error
	return err
}

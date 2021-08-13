package service

import (
	"context"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/domain/repository"
	"github.com/c-4u/company-service/infrastructure/external/topic"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateCompany(ctx context.Context, corporateName, tradeName, cnpj string) (*string, error) {
	company, err := entity.NewCompany(corporateName, tradeName, cnpj)
	if err != nil {
		return nil, err
	}

	err = s.Repository.CreateCompany(ctx, company)
	if err != nil {
		return nil, err
	}

	event, err := entity.NewCompanyEvent(company)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.NEW_COMPANY, company.ID)
	if err != nil {
		return nil, err
	}

	return &company.ID, nil
}

func (s *Service) FindCompany(ctx context.Context, id string) (*entity.Company, error) {
	company, err := s.Repository.FindCompany(ctx, id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (s *Service) SearchCompanies(ctx context.Context, corporateName, tradeName, cnpj string, pageSize int, pageToken string) (*string, []*entity.Company, error) {
	filter, err := entity.NewFilter(corporateName, tradeName, cnpj, pageSize, pageToken)
	if err != nil {
		return nil, nil, err
	}

	nextPageToken, companies, err := s.Repository.SearchCompanies(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	return nextPageToken, companies, nil
}

func (s *Service) UpdateCompany(ctx context.Context, id, corporateName, tradeName string) error {
	company, err := s.Repository.FindCompany(ctx, id)
	if err != nil {
		return err
	}

	if err = company.SetCorporateName(corporateName); err != nil {
		return err
	}
	if err = company.SetTradeName(tradeName); err != nil {
		return err
	}
	if err = s.Repository.SaveCompany(ctx, company); err != nil {
		return err
	}

	event, err := entity.NewCompanyEvent(company)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	if err = s.Repository.PublishEvent(ctx, string(msg), topic.UPDATE_COMPANY, company.ID); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateEmployee(ctx context.Context, id, companyID string) error {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	employee, err := entity.NewEmployee(id, company)
	if err != nil {
		return err
	}

	err = s.Repository.CreateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

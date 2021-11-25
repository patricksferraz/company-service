package service

import (
	"context"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/domain/entity/event"
	"github.com/c-4u/company-service/domain/entity/filter"
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

	event, err := event.NewCompanyEvent(company)
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
	filter, err := filter.NewCompanyFilter(corporateName, tradeName, cnpj, pageSize, pageToken)
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

	event, err := event.NewCompanyEvent(company)
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

func (s *Service) CreateEmployee(ctx context.Context, id string) error {
	employee, err := entity.NewEmployee(id)
	if err != nil {
		return err
	}

	err = s.Repository.CreateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddEmployeeToCompany(ctx context.Context, companyID, employeeID string) error {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	employee, err := s.Repository.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	companyEmployee, err := entity.NewCompanyEmployee(company.ID, employee.ID)
	if err != nil {
		return err
	}

	err = s.Repository.AddEmployeeToCompany(ctx, companyEmployee)
	if err != nil {
		return err
	}

	event, err := event.NewCompanyEmployeeEvent(company.ID, employee.ID)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.ADD_EMPLOYEE_TO_COMPANY, company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateWorkScale(ctx context.Context, name, description, companyID string) (*string, error) {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	workScale, err := entity.NewWorkScale(name, description, company)
	if err != nil {
		return nil, err
	}

	err = s.Repository.CreateWorkScale(ctx, workScale)
	if err != nil {
		return nil, err
	}

	event, err := event.NewWorkScaleEvent(workScale)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.NEW_WORK_SCALE, company.ID)
	if err != nil {
		return nil, err
	}

	return &workScale.ID, nil
}

func (s *Service) FindWorkScale(ctx context.Context, companyID, workScaleID string) (*entity.WorkScale, error) {
	workScale, err := s.Repository.FindWorkScale(ctx, companyID, workScaleID)
	if err != nil {
		return nil, err
	}
	return workScale, nil
}

func (s *Service) SearchWorkScales(ctx context.Context, name, companyID string) ([]*entity.WorkScale, error) {
	filter, err := filter.NewWorkScaleFilter(name, companyID)
	if err != nil {
		return nil, err
	}

	workScales, err := s.Repository.SearchWorkScales(ctx, filter)
	if err != nil {
		return nil, err
	}

	return workScales, nil
}

func (s *Service) AddClockToWorkScale(ctx context.Context, clockType int, clock, timezone, companyID, workScaleID string) (*string, error) {
	workScale, err := s.Repository.FindWorkScale(ctx, companyID, workScaleID)
	if err != nil {
		return nil, err
	}

	c, err := entity.NewClock(clock, clockType, timezone, workScale)
	if err != nil {
		return nil, err
	}

	err = s.Repository.CreateClock(ctx, c)
	if err != nil {
		return nil, err
	}

	event, err := event.NewClockEvent(c)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.ADD_CLOCK_TO_WORK_SCALE, workScale.ID)
	if err != nil {
		return nil, err
	}

	return &c.ID, nil
}

// TODO: use companyID or remove
func (s *Service) FindClock(ctx context.Context, companyID, workScaleID, clockID string) (*entity.Clock, error) {
	clock, err := s.Repository.FindClock(ctx, workScaleID, clockID)
	if err != nil {
		return nil, err
	}
	return clock, nil
}

// TODO: use companyID or remove
func (s *Service) DeleteClock(ctx context.Context, companyID, workScaleID, clockID string) error {
	err := s.Repository.DeleteClock(ctx, workScaleID, clockID)
	if err != nil {
		return err
	}

	event, err := event.NewDeleteClockEvent(companyID, workScaleID, clockID)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	if err = s.Repository.PublishEvent(ctx, string(msg), topic.DELETE_CLOCK, clockID); err != nil {
		return err
	}

	return nil
}

// TODO: use companyID or remove
func (s *Service) UpdateClock(ctx context.Context, clockType int, clock, timezone, companyID, workScaleID, clockID string) error {
	c, err := s.Repository.FindClock(ctx, workScaleID, clockID)
	if err != nil {
		return err
	}

	if err = c.SetType(clockType); err != nil {
		return err
	}
	if err = c.SetClock(clock); err != nil {
		return err
	}
	if err = c.SetTimezone(timezone); err != nil {
		return err
	}
	if err = s.Repository.SaveClock(ctx, c); err != nil {
		return err
	}

	event, err := event.NewClockEvent(c)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	if err = s.Repository.PublishEvent(ctx, string(msg), topic.UPDATE_CLOCK, clockID); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddWorkScaleToEmployee(ctx context.Context, companyID, employeeID, workScaleID string) error {
	companyEmployee, err := s.Repository.FindCompanyEmployee(ctx, companyID, employeeID)
	if err != nil {
		return err
	}

	workScale, err := s.Repository.FindWorkScale(ctx, companyID, workScaleID)
	if err != nil {
		return err
	}

	err = companyEmployee.SetScale(workScale)
	if err != nil {
		return err
	}

	err = s.Repository.SaveCompanyEmployee(ctx, companyEmployee)
	if err != nil {
		return err
	}

	event, err := event.NewWorkScaleEmployeeEvent(companyID, employeeID, workScaleID)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.ADD_WORK_SCALE_TO_EMPLOYEE, workScale.ID)
	if err != nil {
		return err
	}

	return nil
}

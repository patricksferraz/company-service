package repository

import (
	"context"

	"github.com/c-4u/company-service/domain/entity"
	"github.com/c-4u/company-service/domain/entity/filter"
)

type RepositoryInterface interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	FindCompany(ctx context.Context, id string) (*entity.Company, error)
	SearchCompanies(ctx context.Context, companyFilter *filter.CompanyFilter) (*string, []*entity.Company, error)
	SaveCompany(ctx context.Context, employee *entity.Company) error

	PublishEvent(ctx context.Context, msg, topic, key string) error

	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	AddEmployeeToCompany(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error
	FindCompanyEmployee(ctx context.Context, companyID, employeeID string) (*entity.CompaniesEmployee, error)
	SaveCompanyEmployee(ctx context.Context, companyEmployee *entity.CompaniesEmployee) error

	CreateWorkScale(ctx context.Context, workScale *entity.WorkScale) error
	FindWorkScale(ctx context.Context, companyID, workScaleID string) (*entity.WorkScale, error)
	SearchWorkScales(ctx context.Context, workScaleFilter *filter.WorkScaleFilter) ([]*entity.WorkScale, error)
	SaveWorkScale(ctx context.Context, workScale *entity.WorkScale) error

	CreateClock(ctx context.Context, clock *entity.Clock) error
	FindClock(ctx context.Context, workScaleID, clockID string) (*entity.Clock, error)
	DeleteClock(ctx context.Context, workScaleID, clockID string) error
	SaveClock(ctx context.Context, clock *entity.Clock) error
}

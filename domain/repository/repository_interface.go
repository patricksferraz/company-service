package repository

import (
	"context"

	"github.com/c-4u/company-service/domain/entity"
)

type RepositoryInterface interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	FindCompany(ctx context.Context, id string) (*entity.Company, error)
	SearchCompanies(ctx context.Context, filter *entity.Filter) (*string, []*entity.Company, error)
	SaveCompany(ctx context.Context, employee *entity.Company) error

	PublishEvent(ctx context.Context, msg, topic, key string) error

	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error
}

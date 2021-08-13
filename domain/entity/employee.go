package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base      `json:",inline" valid:"required"`
	CompanyID string   `json:"company_id" gorm:"column:company_id;type:uuid;not null" valid:"uuid"`
	Company   *Company `json:"-" valid:"-"`
}

func NewEmployee(id string, company *Company) (*Employee, error) {
	employee := &Employee{
		CompanyID: company.ID,
		Company:   company,
	}
	employee.ID = id
	employee.CreatedAt = time.Now()

	if err := employee.isValid(); err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

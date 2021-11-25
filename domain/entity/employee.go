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
	Companies []*Company `json:"-" gorm:"many2many:companies_employees" valid:"-"`
}

func NewEmployee(id string) (*Employee, error) {
	employee := &Employee{}
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

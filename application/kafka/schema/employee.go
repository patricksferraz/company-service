package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base      `json:",inline" valid:"required"`
	CompanyID string `json:"company_id,omitempty" valid:"uuid"`
}

func NewEmployee() *Employee {
	return &Employee{}
}

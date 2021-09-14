package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base `json:",inline" valid:"required"`
}

func NewEmployee() *Employee {
	return &Employee{}
}

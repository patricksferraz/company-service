package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CompanyEmployeeEvent struct {
	ID         string `json:"id,omitempty" valid:"uuid"`
	CompanyID  string `json:"company_id,omitempty" valid:"uuid"`
	EmployeeID string `json:"employee_id,omitempty" valid:"uuid"`
}

func NewCompanyEmployeeEvent(companyID, employeeID string) (*CompanyEmployeeEvent, error) {
	e := &CompanyEmployeeEvent{
		ID:         uuid.NewV4().String(),
		CompanyID:  companyID,
		EmployeeID: employeeID,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *CompanyEmployeeEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CompanyEmployeeEvent) ToJson() ([]byte, error) {
	err := e.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(e)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

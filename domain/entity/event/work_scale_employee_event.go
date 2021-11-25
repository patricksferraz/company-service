package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkScaleEmployeeEvent struct {
	ID          string `json:"id" valid:"uuid"`
	CompanyID   string `json:"company_id" valid:"uuid"`
	EmployeeID  string `json:"employee_id" valid:"uuid"`
	WorkScaleID string `json:"work_scale_id" valid:"uuid"`
}

func NewWorkScaleEmployeeEvent(companyID, employeeID, workScaleID string) (*WorkScaleEmployeeEvent, error) {
	e := &WorkScaleEmployeeEvent{
		ID:          uuid.NewV4().String(),
		CompanyID:   companyID,
		EmployeeID:  employeeID,
		WorkScaleID: workScaleID,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *WorkScaleEmployeeEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *WorkScaleEmployeeEvent) ToJson() ([]byte, error) {
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

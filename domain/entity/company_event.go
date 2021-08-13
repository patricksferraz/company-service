package entity

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type CompanyEvent struct {
	ID      string   `json:"id,omitempty" valid:"uuid"`
	Company *Company `json:"company,omitempty" valid:"-"`
}

func NewCompanyEvent(company *Company) (*CompanyEvent, error) {
	e := &CompanyEvent{
		ID:      uuid.NewV4().String(),
		Company: company,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *CompanyEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CompanyEvent) ToJson() ([]byte, error) {
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

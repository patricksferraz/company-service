package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/company-service/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type ClockEvent struct {
	ID    string        `json:"id,omitempty" valid:"uuid"`
	Clock *entity.Clock `json:"clock,omitempty" valid:"-"`
}

func NewClockEvent(clock *entity.Clock) (*ClockEvent, error) {
	e := &ClockEvent{
		ID:    uuid.NewV4().String(),
		Clock: clock,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *ClockEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *ClockEvent) ToJson() ([]byte, error) {
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

type DeleteClockEvent struct {
	ID          string `json:"id,omitempty" valid:"uuid"`
	CompanyID   string `json:"company_id,omitempty" valid:"uuid"`
	WorkScaleID string `json:"work_scale_id,omitempty" valid:"uuid"`
	ClockID     string `json:"clock_id,omitempty" valid:"uuid"`
}

func NewDeleteClockEvent(companyID, workScaleID, clockID string) (*DeleteClockEvent, error) {
	e := &DeleteClockEvent{
		ID:          uuid.NewV4().String(),
		CompanyID:   companyID,
		WorkScaleID: workScaleID,
		ClockID:     clockID,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *DeleteClockEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *DeleteClockEvent) ToJson() ([]byte, error) {
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

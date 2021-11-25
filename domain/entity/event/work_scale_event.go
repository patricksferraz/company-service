package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/company-service/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type WorkScaleEvent struct {
	ID        string            `json:"id,omitempty" valid:"uuid"`
	WorkScale *entity.WorkScale `json:"work_scale,omitempty" valid:"-"`
}

func NewWorkScaleEvent(workScale *entity.WorkScale) (*WorkScaleEvent, error) {
	e := &WorkScaleEvent{
		ID:        uuid.NewV4().String(),
		WorkScale: workScale,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *WorkScaleEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *WorkScaleEvent) ToJson() ([]byte, error) {
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

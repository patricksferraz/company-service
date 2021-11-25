package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkScale struct {
	Base        `json:",inline" valid:"required"`
	Name        string   `json:"name" gorm:"column:name;not null" valid:"required"`
	Description string   `json:"description,omitempty" gorm:"column:description" valid:"-"`
	Clocks      []*Clock `json:"clocks,omitempty" gorm:"ForeignKey:WorkScaleID" valid:"-"`
	CompanyID   *string  `json:"company_id,omitempty" gorm:"column:company_id;type:uuid;not null" valid:"uuid"`
	Company     *Company `json:"-" valid:"-"`
}

func NewWorkScale(name string, description string, company *Company, clocks ...*Clock) (*WorkScale, error) {
	entity := &WorkScale{
		Name:        name,
		Description: description,
		Clocks:      clocks,
		CompanyID:   &company.ID,
		Company:     company,
	}
	entity.ID = uuid.NewV4().String()
	entity.CreatedAt = time.Now()

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *WorkScale) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

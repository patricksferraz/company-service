package filter

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkScaleFilter struct {
	Name      string `json:"name" valid:"optional"`
	CompanyID string `json:"company_id" valid:"uuid"`
}

func (e *WorkScaleFilter) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewWorkScaleFilter(name, companyID string) (*WorkScaleFilter, error) {
	e := &WorkScaleFilter{
		Name:      name,
		CompanyID: companyID,
	}

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}

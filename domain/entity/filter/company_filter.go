package filter

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CompanyFilter struct {
	CorporateName string `json:"corporate_name" valid:"optional"`
	TradeName     string `json:"trade_name" valid:"optional"`
	Cnpj          string `json:"cnpj" valid:"optional"`
	PageSize      int    `json:"page_size" valid:"optional"`
	PageToken     string `json:"page_token" valid:"optional"`
}

func (e *CompanyFilter) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewCompanyFilter(corporateName, tradeName, cnpj string, pageSize int, pageToken string) (*CompanyFilter, error) {

	if pageSize == 0 {
		pageSize = 10
	}

	e := &CompanyFilter{
		CorporateName: corporateName,
		TradeName:     tradeName,
		Cnpj:          cnpj,
		PageSize:      pageSize,
		PageToken:     pageToken,
	}

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}

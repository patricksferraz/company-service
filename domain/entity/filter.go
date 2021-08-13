package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Filter struct {
	CorporateName string `json:"corporate_name" valid:"optional"`
	TradeName     string `json:"trade_name" valid:"optional"`
	Cnpj          string `json:"cnpj" valid:"optional"`
	PageSize      int    `json:"page_size" valid:"optional"`
	PageToken     string `json:"page_token" valid:"optional"`
}

func (e *Filter) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewFilter(corporateName, tradeName, cnpj string, pageSize int, pageToken string) (*Filter, error) {

	if pageSize == 0 {
		pageSize = 10
	}

	e := &Filter{
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

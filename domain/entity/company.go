package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/company-service/utils"
	"github.com/paemuri/brdoc"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.TagMap["cnpj"] = govalidator.Validator(func(str string) bool {
		return brdoc.IsCNPJ(str)
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type Company struct {
	Base          `json:",inline" valid:"required"`
	CorporateName string      `json:"corporate_name,omitempty" gorm:"column:corporate_name;type:varchar(255);not null" valid:"required"`
	TradeName     string      `json:"trade_name,omitempty" gorm:"column:trade_name;type:varchar(255);not null" valid:"required"`
	Cnpj          string      `json:"cnpj,omitempty" gorm:"column:cnpj;type:varchar(25);not null;unique" valid:"cnpj"`
	Token         *string     `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	Employees     []*Employee `json:"employees,omitempty" gorm:"ForeignKey:CompanyID" valid:"-"`
}

func NewCompany(corporateName, tradeName, cnpj string) (*Company, error) {

	utils.CleanNonDigits(&cnpj)
	token := primitive.NewObjectID().Hex()
	company := &Company{
		CorporateName: corporateName,
		TradeName:     tradeName,
		Cnpj:          cnpj,
		Token:         &token,
	}
	company.ID = uuid.NewV4().String()
	company.CreatedAt = time.Now()

	if err := company.isValid(); err != nil {
		return nil, err
	}

	return company, nil
}

func (e *Company) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Company) SetCorporateName(corporateName string) error {
	e.CorporateName = corporateName
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Company) SetTradeName(tradeName string) error {
	e.TradeName = tradeName
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

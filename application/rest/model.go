package rest

import "time"

type Base struct {
	ID        string    `json:"id,omitempty" binding:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty" time_format:"RFC3339"`
	UpdatedAt time.Time `json:"updated_at,omitempty" time_format:"RFC3339"`
}

type Company struct {
	Base          `json:",inline"`
	CorporateName string `json:"corporate_name,omitempty"`
	TradeName     string `json:"trade_name,omitempty"`
	Cnpj          string `json:"cnpj,omitempty"`
}

type CreateCompanyRequest struct {
	CorporateName string `json:"corporate_name" binding:"required"`
	TradeName     string `json:"trade_name" binding:"required"`
	Cnpj          string `json:"cnpj" binding:"required"`
}

type CreateCompanyResponse struct {
	ID string `json:"id"`
}

type IDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type HTTPResponse struct {
	Code    int    `json:"code,omitempty" example:"200"`
	Message string `json:"message,omitempty" example:"a message"`
}

type HTTPError struct {
	Code  int    `json:"code,omitempty" example:"400"`
	Error string `json:"error,omitempty" example:"status bad request"`
}

type SearchCompaniesRequest struct {
	Filter `json:",inline"`
}

type Filter struct {
	CorporateName string `json:"corporate_name" form:"corporate_name"`
	TradeName     string `json:"trade_name" form:"trade_name"`
	Cnpj          string `json:"cnpj" form:"cnpj"`
	PageSize      int    `json:"page_size" form:"page_size" default:"10"`
	PageToken     string `json:"page_token" form:"page_token"`
}

type SearchCompaniesResponse struct {
	NextPageToken string    `json:"next_page_token"`
	Employees     []Company `json:"companies"`
}

type UpdateCompanyRequest struct {
	CorporateName string `json:"corporate_name"`
	TradeName     string `json:"trade_name"`
}

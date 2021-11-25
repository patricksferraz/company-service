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
	CompanyID string `uri:"company_id" binding:"required,uuid"`
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
	CompanyFilter `json:",inline"`
}

type CompanyFilter struct {
	CorporateName string `json:"corporate_name" form:"corporate_name"`
	TradeName     string `json:"trade_name" form:"trade_name"`
	Cnpj          string `json:"cnpj" form:"cnpj"`
	PageSize      int    `json:"page_size" form:"page_size" default:"10"`
	PageToken     string `json:"page_token" form:"page_token"`
}

type SearchCompaniesResponse struct {
	NextPageToken string    `json:"next_page_token"`
	Companies     []Company `json:"companies"`
}

type UpdateCompanyRequest struct {
	CorporateName string `json:"corporate_name"`
	TradeName     string `json:"trade_name"`
}

type AddEmployeeToCompanyRequest struct {
	EmployeeID string `uri:"employee_id" binding:"required,uuid"`
	CompanyID  string `uri:"company_id" binding:"required,uuid"`
}

type Clock struct {
	Base        `json:",inline"`
	Type        string `json:"type,omitempty"`
	Clock       string `json:"clock,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	WorkScaleID string `json:"work_scale_id,omitempty"`
}

type WorkScale struct {
	Base        `json:",inline"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	CompanyID   string  `json:"company_id,omitempty"`
	Clocks      []Clock `json:"clocks,omitempty"`
}

type CreateWorkScaleRequest struct {
	Name        string `json:"name,omitempty" binding:"required"`
	Description string `json:"description"`
}

type CreateWorkScaleResponse struct {
	ID string `json:"id"`
}

type FindWorkScaleRequest struct {
	CompanyID   string `uri:"company_id" binding:"required,uuid"`
	WorkScaleID string `uri:"work_scale_id" binding:"required,uuid"`
}

type SearchWorkScalesRequest struct {
	WorkScaleFilter `json:",inline"`
}

type WorkScaleFilter struct {
	Name string `json:"name" form:"name"`
}

type SearchWorkScalesResponse struct {
	WorkScales []WorkScale `json:"work_scales"`
}

type AddClockToWorkScaleIDRequest struct {
	CompanyID   string `uri:"company_id" binding:"required,uuid"`
	WorkScaleID string `uri:"work_scale_id" binding:"required,uuid"`
}

type AddClockToWorkScaleRequest struct {
	Type     int    `json:"type" binding:"required"`
	Clock    string `json:"clock" binding:"required"`
	Timezone string `json:"timezone" binding:"required"`
}

type AddClockToWorkScaleResponse struct {
	ID string `json:"id"`
}

type FindClockRequest struct {
	CompanyID   string `uri:"company_id" binding:"required,uuid"`
	WorkScaleID string `uri:"work_scale_id" binding:"required,uuid"`
	ClockID     string `uri:"clock_id" binding:"required,uuid"`
}

type DeleteClockRequest struct {
	CompanyID   string `uri:"company_id" binding:"required,uuid"`
	WorkScaleID string `uri:"work_scale_id" binding:"required,uuid"`
	ClockID     string `uri:"clock_id" binding:"required,uuid"`
}

type UpdateClockIDRequest struct {
	CompanyID   string `uri:"company_id" binding:"required,uuid"`
	WorkScaleID string `uri:"work_scale_id" binding:"required,uuid"`
	ClockID     string `uri:"clock_id" binding:"required,uuid"`
}

type UpdateClockRequest struct {
	Type     int    `json:"type"`
	Clock    string `json:"clock"`
	Timezone string `json:"timezone"`
}

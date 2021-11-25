package rest

import (
	"net/http"

	"github.com/c-4u/company-service/domain/service"
	"github.com/gin-gonic/gin"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateCompany godoc
// @Security ApiKeyAuth
// @Summary create a company
// @Description create company
// @ID createCompany
// @Tags Company
// @Accept json
// @Produce json
// @Param body body CreateCompanyRequest true "JSON body to create a new company"
// @Success 200 {object} CreateCompanyResponse
// @Failure 401 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies [post]
func (s *RestService) CreateCompany(ctx *gin.Context) {
	var json CreateCompanyRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	id, err := s.Service.CreateCompany(ctx, json.CorporateName, json.TradeName, json.Cnpj)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, CreateCompanyResponse{ID: *id})
}

// FindCompany godoc
// @Security ApiKeyAuth
// @Summary find a company
// @Description Router for find a company
// @ID findCompany
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Success 200 {object} Company
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id} [get]
func (s *RestService) FindCompany(ctx *gin.Context) {
	var req IDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	company, err := s.Service.FindCompany(ctx, req.CompanyID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// SearchCompanies godoc
// @Security ApiKeyAuth
// @Summary search companies by filter
// @ID searchCompanies
// @Tags Company
// @Description Search companies by `filter`. if the page size is empty, 10 will be considered.
// @Accept json
// @Produce json
// @Param body query SearchCompaniesRequest true "JSON body for search companies"
// @Success 200 {object} SearchCompaniesResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies [get]
func (s *RestService) SearchCompanies(ctx *gin.Context) {
	var body SearchCompaniesRequest

	if err := ctx.ShouldBindQuery(&body); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	nextPageToken, companies, err := s.Service.SearchCompanies(ctx, body.CorporateName, body.TradeName, body.Cnpj, body.PageSize, body.PageToken)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"next_page_token": *nextPageToken,
			"companies":       companies,
		},
	)
}

// UpdateCompany godoc
// @Security ApiKeyAuth
// @Summary update a Company
// @Description Router for update a company
// @ID updateCompany
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param body body UpdateCompanyRequest true "JSON body to update a new company"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id} [put]
func (s *RestService) UpdateCompany(ctx *gin.Context) {
	var req IDRequest
	var json UpdateCompanyRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	err := s.Service.UpdateCompany(ctx, req.CompanyID, json.CorporateName, json.TradeName)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, HTTPResponse{Code: http.StatusOK, Message: "updated successfully"})
}

// AddEmployeeToCompany godoc
// @Security ApiKeyAuth
// @Summary add employee to company
// @Description Router for add employee to company
// @ID addEmployeeToCompany
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/employees/{employee_id} [post]
func (s *RestService) AddEmployeeToCompany(ctx *gin.Context) {
	var req AddEmployeeToCompanyRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	err := s.Service.AddEmployeeToCompany(ctx, req.CompanyID, req.EmployeeID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, HTTPResponse{Code: http.StatusOK, Message: "added successfully"})
}

// CreateWorkScale godoc
// @Security ApiKeyAuth
// @Summary create work scale
// @Description Create work scale
// @ID createWorkScale
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param body body CreateWorkScaleRequest true "JSON body to create a new work scale"
// @Success 200 {object} CreateWorkScaleResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales [post]
func (s *RestService) CreateWorkScale(ctx *gin.Context) {
	var req IDRequest
	var json CreateWorkScaleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	id, err := s.Service.CreateWorkScale(ctx, json.Name, json.Description, req.CompanyID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, CreateWorkScaleResponse{ID: *id})
}

// FindWorkScale godoc
// @Security ApiKeyAuth
// @Summary find a work scale
// @Description Router for find a work scale
// @ID findWorkScale
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param work_scale_id path string true "Work Scale ID"
// @Success 200 {object} WorkScale
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales/{work_scale_id} [get]
func (s *RestService) FindWorkScale(ctx *gin.Context) {
	var req FindWorkScaleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	company, err := s.Service.FindWorkScale(ctx, req.CompanyID, req.WorkScaleID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// SearchWorkScales godoc
// @Security ApiKeyAuth
// @Summary search work scales by filter
// @ID searchWorkScales
// @Tags Company
// @Description Search work scales by `filter`.
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param body query SearchWorkScalesRequest true "JSON body for search work scales"
// @Success 200 {object} SearchWorkScalesResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales [get]
func (s *RestService) SearchWorkScales(ctx *gin.Context) {
	var req IDRequest
	var body SearchWorkScalesRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindQuery(&body); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	workScales, err := s.Service.SearchWorkScales(ctx, body.Name, req.CompanyID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"work_scales": workScales,
		},
	)
}

// AddClockToWorkScale godoc
// @Security ApiKeyAuth
// @Summary add clock to work scale
// @Description Add clock to work scale
// @ID addClockToWorkScale
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param work_scale_id path string true "Work Scale ID"
// @Param body body AddClockToWorkScaleRequest true "JSON body to add clock to work scale"
// @Success 200 {object} AddClockToWorkScaleResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales/{work_scale_id}/clocks [post]
func (s *RestService) AddTimeToWorkScale(ctx *gin.Context) {
	var req AddClockToWorkScaleIDRequest
	var json AddClockToWorkScaleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	id, err := s.Service.AddClockToWorkScale(ctx, json.Type, json.Clock, json.Timezone, req.CompanyID, req.WorkScaleID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, CreateWorkScaleResponse{ID: *id})
}

// FindClock godoc
// @Security ApiKeyAuth
// @Summary find clock
// @Description Find clock
// @ID findClock
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param work_scale_id path string true "Work Scale ID"
// @Param clock_id path string true "Clock ID"
// @Success 200 {object} Clock
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales/{work_scale_id}/clocks/{clock_id} [get]
func (s *RestService) FindClock(ctx *gin.Context) {
	var req FindClockRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	res, err := s.Service.FindClock(ctx, req.CompanyID, req.WorkScaleID, req.ClockID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteClock godoc
// @Security ApiKeyAuth
// @Summary delete clock
// @Description Router for delete clock
// @ID deleteClock
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param work_scale_id path string true "Work Scale ID"
// @Param clock_id path string true "Clock ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales/{work_scale_id}/clocks/{clock_id} [delete]
func (s *RestService) DeleteClock(ctx *gin.Context) {
	var req DeleteClockRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	err := s.Service.DeleteClock(ctx, req.CompanyID, req.WorkScaleID, req.ClockID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, HTTPResponse{Code: http.StatusOK, Message: "deleted successfully"})
}

// UpdateClock godoc
// @Security ApiKeyAuth
// @Summary update a clock
// @Description Router for update a clock
// @ID updateClock
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param work_scale_id path string true "Work Scale ID"
// @Param clock_id path string true "Clock ID"
// @Param body body UpdateClockRequest true "JSON body to update a clock"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{company_id}/work-scales/{work_scale_id}/clocks/{clock_id} [put]
func (s *RestService) UpdateClock(ctx *gin.Context) {
	var req UpdateClockIDRequest
	var json UpdateClockRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	err := s.Service.UpdateClock(ctx, json.Type, json.Clock, json.Timezone, req.CompanyID, req.WorkScaleID, req.ClockID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, HTTPResponse{Code: http.StatusOK, Message: "updated successfully"})
}

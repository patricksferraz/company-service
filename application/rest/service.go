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
// @Param id path string true "Company ID"
// @Success 200 {object} Company
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{id} [get]
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

	company, err := s.Service.FindCompany(ctx, req.ID)
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
// @Success 200 {array} SearchCompaniesResponse
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
// @Param id path string true "Company ID"
// @Param body body UpdateCompanyRequest true "JSON body to update a new company"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /companies/{id} [put]
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

	err := s.Service.UpdateCompany(ctx, req.ID, json.CorporateName, json.TradeName)
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
// @Router /companies/{company_id}/employee/{employee_id} [post]
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

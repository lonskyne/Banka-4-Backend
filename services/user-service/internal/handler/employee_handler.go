package handler

import (
	"net/http"
	"strconv"

	"common/pkg/errors"
	"user-service/internal/dto"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

// ZaposlenHandler rukuje HTTP zahtevima za zaposlene
type EmployeeHandler struct {
	service *service.EmployeeService
}

// NoviZaposlenHandler kreira novi EmployeeHandler
func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// DohvatiListuZaposlenih dohvata listu svih zaposlenih sa filtriranjem i paginacijom
// GET /employees?email=petar&firstName=Petar&lastName=Petrović&position=Menadžer&page=1&pageSize=10
func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
	var query dto.ListEmployeesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Error(errors.BadRequestErr(err.Error()))
		return
	}

	// Postavi default vrednosti
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}

	result, err := h.service.GetAllEmployees(c.Request.Context(), &query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DohvatiZaposlenog dohvata jednog zaposlenog po ID-u
// GET /employees/:id
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(errors.BadRequestErr("nevalidan ID zaposlenog"))
		return
	}

	employee, err := h.service.GetEmployee(c.Request.Context(), uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

// KreirajZaposlenog kreira novog zaposlenog
// POST /employees
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequestErr(err.Error()))
		return
	}

	employee, err := h.service.CreateEmployee(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// AzurirajZaposlenog ažurira postojećeg zaposlenog
// PUT /employees/:id
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(errors.BadRequestErr("nevalidan ID zaposlenog"))
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequestErr(err.Error()))
		return
	}

	employee, err := h.service.UpdateEmployee(c.Request.Context(), uint(id), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

// DeaktivujZaposlenog deaktivira zaposlenog
// DELETE /employees/:id
func (h *EmployeeHandler) DeactivateEmployee(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(errors.BadRequestErr("nevalidan ID zaposlenog"))
		return
	}

	if err := h.service.DeactivateEmployee(c.Request.Context(), uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "zaposleni je uspešno deaktiviran"})
}

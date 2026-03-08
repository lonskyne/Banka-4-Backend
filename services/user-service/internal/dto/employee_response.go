package dto

import (
	"time"
	"user-service/internal/model"
)

// OdgovorsZaposlen predstavlja odgovor sa podacima zaposlenog (bez osetljivih polja)
type EmployeeResponse struct {
	EmployeeID  uint      `json:"employeeId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Username    string    `json:"username"`
	Department  string    `json:"department"`
	PositionID  uint      `json:"positionId"`
	Position    string    `json:"position"`
	Active      bool      `json:"active"`
}

// OdgovorListe obavija listu zaposlenih sa informacijama o paginaciji
type ListEmployeesResponse struct {
	Data       []EmployeeResponse `json:"data"`
	Total      int64              `json:"total"`
	Page       int                `json:"page"`
	PageSize   int                `json:"pageSize"`
	TotalPages int                `json:"totalPages"`
}

// KonvertujUOdgovor pretvara model.Employee u EmployeeResponse
func ToEmployeeResponse(e *model.Employee) *EmployeeResponse {
	positionName := ""
	if e.Position.PositionID > 0 {
		positionName = e.Position.Title
	}

	return &EmployeeResponse{
		EmployeeID:  e.EmployeeID,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		DateOfBirth: e.DateOfBirth,
		Gender:      e.Gender,
		Email:       e.Email,
		PhoneNumber: e.PhoneNumber,
		Address:     e.Address,
		Username:    e.Username,
		Department:  e.Department,
		PositionID:  e.PositionID,
		Position:    positionName,
		Active:      e.Active,
	}
}

// KonvertujListuUOdgovor pretvara niz model.Employee u ListEmployeesResponse
func ToEmployeeResponseList(employees []model.Employee, total int64, page, pageSize int) *ListEmployeesResponse {
	responses := make([]EmployeeResponse, len(employees))
	for i, e := range employees {
		responses[i] = *ToEmployeeResponse(&e)
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return &ListEmployeesResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

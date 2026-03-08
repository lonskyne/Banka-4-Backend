package dto

import "time"

// ZahtevZaKreiranje predstavlja podatke koji se šalju pri kreiranju novog zaposlenog
type CreateEmployeeRequest struct {
	FirstName   string    `json:"firstName" binding:"required,min=1"`
	LastName    string    `json:"lastName" binding:"required,min=1"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email" binding:"required,email"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Username    string    `json:"username" binding:"required,min=3"`
	Password    string    `json:"password" binding:"required,min=8"`
	Active      bool      `json:"active"`
	Department  string    `json:"department"`
	PositionID  uint      `json:"positionId"`
}

// ZahtevZaAzuriranje predstavlja podatke za ažuriranje zaposlenog
type UpdateEmployeeRequest struct {
	LastName    *string `json:"lastName"`
	Gender      *string `json:"gender"`
	PhoneNumber *string `json:"phoneNumber"`
	Address     *string `json:"address"`
	Department  *string `json:"department"`
	PositionID  *uint   `json:"positionId"`
	Active      *bool   `json:"active"`
}

// UpitiZaListu predstavlja parametre za filtriranje zaposlenih
type ListEmployeesQuery struct {
	Email     string `form:"email"`
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Position  string `form:"position"`
	Page      int    `form:"page,default=1" binding:"min=1"`
	PageSize  int    `form:"pageSize,default=10" binding:"min=1,max=100"`
}

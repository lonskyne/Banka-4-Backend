package repository

import (
	"context"
	"errors"

	"user-service/internal/model"

	"gorm.io/gorm"
)

// ZaposlenRepozitorijum upravlja bazom podataka za zaposlene
type EmployeeRepository struct {
	db *gorm.DB
}

// NoviZaposlenRepozitorijum kreira novi EmployeeRepository
func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

// DohvatiSve dohvata sve zaposlene sa opcionalnim filtriranjem i paginacijom
func (r *EmployeeRepository) GetAll(ctx context.Context, email, firstName, lastName, position string, page, pageSize int) ([]model.Employee, int64, error) {
	var employees []model.Employee
	var total int64

	query := r.db.WithContext(ctx).Preload("Position")

	// Primeni filtere
	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}
	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+lastName+"%")
	}
	if position != "" {
		query = query.Where("position_name ILIKE ?", "%"+position+"%")
	}

	// Dobij ukupan broj
	if err := query.Model(&model.Employee{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Primeni paginaciju
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("employee_id DESC").Find(&employees).Error; err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

// DohvatiPoID dohvata zaposlenog po ID-u
func (r *EmployeeRepository) GetByID(ctx context.Context, id uint) (*model.Employee, error) {
	var employee model.Employee
	result := r.db.WithContext(ctx).Preload("Position").First(&employee, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &employee, result.Error
}

// DohvatiPoEmail dohvata zaposlenog po emailu
func (r *EmployeeRepository) GetByEmail(ctx context.Context, email string) (*model.Employee, error) {
	var employee model.Employee
	result := r.db.WithContext(ctx).Preload("Position").Where("email = ?", email).First(&employee)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &employee, result.Error
}

// DohvatiPoUsername dohvata zaposlenog po korisničkom imenu
func (r *EmployeeRepository) GetByUsername(ctx context.Context, username string) (*model.Employee, error) {
	var employee model.Employee
	result := r.db.WithContext(ctx).Preload("Position").Where("username = ?", username).First(&employee)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &employee, result.Error
}

// Kreiraj kreira novog zaposlenog
func (r *EmployeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

// Azuriraj ažurira zaposlenog
func (r *EmployeeRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) (*model.Employee, error) {
	result := r.db.WithContext(ctx).Model(&model.Employee{}).Where("employee_id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.GetByID(ctx, id)
}

// Obriši deaktivira zaposlenog
func (r *EmployeeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.Employee{}).Where("employee_id = ?", id).Update("active", false).Error
}

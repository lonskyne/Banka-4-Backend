package service

import (
	"context"

	"common/pkg/errors"
	"user-service/internal/dto"
	"user-service/internal/model"
	"user-service/internal/repository"
)

// ZaposlenServis upravlja poslovnom logikom za zaposlene
type EmployeeService struct {
	repo *repository.EmployeeRepository
}

// NoviZaposlenServis kreira novi EmployeeService
func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// DohvatiSveZaposlene dohvata sve zaposlene sa filtriranjem i paginacijom
func (s *EmployeeService) GetAllEmployees(ctx context.Context, query *dto.ListEmployeesQuery) (*dto.ListEmployeesResponse, error) {
	employees, total, err := s.repo.GetAll(ctx, query.Email, query.FirstName, query.LastName, query.Position, query.Page, query.PageSize)
	if err != nil {
		return nil, errors.InternalErr(err)
	}

	return dto.ToEmployeeResponseList(employees, total, query.Page, query.PageSize), nil
}

// DohvatiZaposlenog dohvata jednog zaposlenog po ID-u
func (s *EmployeeService) GetEmployee(ctx context.Context, id uint) (*dto.EmployeeResponse, error) {
	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.InternalErr(err)
	}

	if employee == nil {
		return nil, errors.NotFoundErr("zaposleni nije pronađen")
	}

	return dto.ToEmployeeResponse(employee), nil
}

// KreirajZaposlenog kreira novog zaposlenog
func (s *EmployeeService) CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	// Proveri da li email već postoji
	existing, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.InternalErr(err)
	}
	if existing != nil {
		return nil, errors.ConflictErr("email je već registrovan")
	}

	// Proveri da li korisničko ime već postoji
	existingUsername, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.InternalErr(err)
	}
	if existingUsername != nil {
		return nil, errors.ConflictErr("korisničko ime je već zauzeto")
	}

	employee := &model.Employee{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Username:    req.Username,
		Password:    req.Password,
		Active:      req.Active,
		Department:  req.Department,
		PositionID:  req.PositionID,
	}

	if err := s.repo.Create(ctx, employee); err != nil {
		return nil, errors.InternalErr(err)
	}

	return dto.ToEmployeeResponse(employee), nil
}

// AzurirajZaposlenog ažurira postojećeg zaposlenog
func (s *EmployeeService) UpdateEmployee(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	// Proveri da li zaposleni postoji
	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.InternalErr(err)
	}
	if employee == nil {
		return nil, errors.NotFoundErr("zaposleni nije pronađen")
	}

	// Izgradi mapu ažuriranja (samo uključi poslata polja)
	updates := make(map[string]interface{})
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.PhoneNumber != nil {
		updates["phone_number"] = *req.PhoneNumber
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.Department != nil {
		updates["department"] = *req.Department
	}
	if req.PositionID != nil {
		updates["position_id"] = *req.PositionID
	}
	if req.Active != nil {
		updates["active"] = *req.Active
	}

	if len(updates) == 0 {
		return dto.ToEmployeeResponse(employee), nil
	}

	updated, err := s.repo.Update(ctx, id, updates)
	if err != nil {
		return nil, errors.InternalErr(err)
	}

	return dto.ToEmployeeResponse(updated), nil
}

// DeaktivujZaposlenog deaktivira zaposlenog
func (s *EmployeeService) DeactivateEmployee(ctx context.Context, id uint) error {
	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.InternalErr(err)
	}
	if employee == nil {
		return errors.NotFoundErr("zaposleni nije pronađen")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.InternalErr(err)
	}

	return nil
}

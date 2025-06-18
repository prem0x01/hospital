package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/repository"
	//"github.com/prem0x01/hospital/internal/utils"
)

type PatientService struct {
	patientRepo *repository.PatientRepository
}

func NewPatientService(patientRepo *repository.PatientRepository) *PatientService {
	return &PatientService{patientRepo: patientRepo}
}

func (s *PatientService) GetPatients(limit, offset int) ([]domain.Patient, error) {
	ctx := context.Background()
	return s.patientRepo.GetAll(ctx, int32(limit), int32(offset))
}

func (s *PatientService) GetPatient(id int) (*domain.Patient, error) {
	ctx := context.Background()
	return s.patientRepo.GetByID(ctx, int32(id))
}

func (s *PatientService) CreatePatient(req *domain.CreatePatientRequest, createdBy int) (*domain.Patient, error) {
	createdByInt32 := int32(createdBy)
	patient := &domain.Patient{
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		Email:                 req.Email,
		Phone:                 req.Phone,
		Gender:                req.Gender,
		Address:               req.Address,
		MedicalHistory:        req.MedicalHistory,
		Allergies:             req.Allergies,
		EmergencyContactName:  req.EmergencyContactName,
		EmergencyContactPhone: req.EmergencyContactPhone,
		CreatedBy:             &createdByInt32,
	}

	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, err
		}
		patient.DateOfBirth = pgtype.Date{Time: dob, Valid: true}
	}
	ctx := context.Background()
	if _, err := s.patientRepo.Create(ctx, *patient); err != nil {
		return nil, err
	}

	return patient, nil
}


func (s *PatientService) UpdatePatient(id int, req *domain.UpdatePatientRequest) (*domain.Patient, error) {
	ctx := context.Background()

	existing, err := s.patientRepo.GetByID(ctx, int32(id))
	if err != nil {
		return nil, fmt.Errorf("patient not found: %w", err)
	}

	if req.FirstName != nil {
		existing.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		existing.LastName = *req.LastName
	}
	if req.Email != nil {
		existing.Email = req.Email
	}
	if req.Phone != nil {
		existing.Phone = req.Phone
	}
	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date_of_birth: %w", err)
		}
		var pgDob pgtype.Date
		if err := pgDob.Scan(dob); err != nil {
			return nil, fmt.Errorf("date scan error: %w", err)
		}
		existing.DateOfBirth = pgDob
	}

	if req.Gender != nil {
		existing.Gender = req.Gender
	}
	if req.Address != nil {
		existing.Address = req.Address
	}
	if req.MedicalHistory != nil {
		existing.MedicalHistory = req.MedicalHistory
	}
	if req.Allergies != nil {
		existing.Allergies = req.Allergies
	}
	if req.EmergencyContactName != nil {
		existing.EmergencyContactName = req.EmergencyContactName
	}
	if req.EmergencyContactPhone != nil {
		existing.EmergencyContactPhone = req.EmergencyContactPhone
	}

	return s.patientRepo.Update(ctx, *existing)
}

func (s *PatientService) DeletePatient(id int) error {
	ctx := context.Background()
	return s.patientRepo.Delete(ctx, int32(id))
}

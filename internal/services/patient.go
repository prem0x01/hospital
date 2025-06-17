package services

import (
	"context"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/repository"
	"time"
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
		CreatedBy:             &createdBy,
	}

	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, err
		}
		patient.DateOfBirth = &domain.NullDate{Time: dob, Valid: true}
	}
	ctx := context.Background()
	if _, err := s.patientRepo.Create(ctx, patient); err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *PatientService) UpdatePatient(id int, req *domain.UpdatePatientRequest) error {
	updates := make(map[string]interface{})

	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return err
		}
		updates["date_of_birth"] = dob
	}
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.MedicalHistory != nil {
		updates["medical_history"] = *req.MedicalHistory
	}
	if req.Allergies != nil {
		updates["allergies"] = *req.Allergies
	}
	if req.EmergencyContactName != nil {
		updates["emergency_contact_name"] = *req.EmergencyContactName
	}
	if req.EmergencyContactPhone != nil {
		updates["emergency_contact_phone"] = *req.EmergencyContactPhone
	}
	ctx := context.Background()

	return s.patientRepo.Update(ctx, int32(id), updates)
}

func (s *PatientService) DeletePatient(id int) error {
	ctx := context.Background()
	return s.patientRepo.Delete(ctx, int32(id))
}

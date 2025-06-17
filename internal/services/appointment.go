package services

import (
	"context"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/repository"
	"time"
)

type AppointmentService struct {
	appointmentRepo *repository.AppointmentRepository
	patientRepo     *repository.PatientRepository
}

func NewAppointmentService(appointmentRepo *repository.AppointmentRepository, patientRepo *repository.PatientRepository) *AppointmentService {
	return &AppointmentService{
		appointmentRepo: appointmentRepo,
		patientRepo:     patientRepo,
	}
}

func (s *AppointmentService) GetAppointments(limit, offset int, userRole string, userID int) ([]domain.Appointment, error) {
	ctx := context.Background()
	return s.appointmentRepo.GetAll(ctx, int32(limit), int32(offset), userRole, int32(userID))
}

func (s *AppointmentService) GetAppointment(id int) (*domain.Appointment, error) {
	ctx := context.Background()
	return s.appointmentRepo.GetByID(ctx, int32(id))
}

func (s *AppointmentService) CreateAppointment(req *domain.CreateAppointmentRequest, createdBy int) (*domain.Appointment, error) {
	ctx := context.Background()
	_, err := s.patientRepo.GetByID(ctx, int32(req.PatientID))
	if err != nil {
		return nil, err
	}

	appointmentDate, err := time.Parse("2006-01-02T15:04", req.AppointmentDate)
	if err != nil {
		return nil, err
	}

	appointment := &domain.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: appointmentDate,
		Notes:           req.Notes,
		CreatedBy:       &createdBy,
	}

	if err := s.appointmentRepo.Create(ctx, appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *AppointmentService) UpdateAppointment(id int, req *domain.UpdateAppointmentRequest) error {
	updates := make(map[string]interface{})

	if req.DoctorID != nil {
		updates["doctor_id"] = *req.DoctorID
	}
	if req.AppointmentDate != nil && *req.AppointmentDate != "" {
		appointmentDate, err := time.Parse("2006-01-02T15:04", *req.AppointmentDate)
		if err != nil {
			return err
		}
		updates["appointment_date"] = appointmentDate
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Notes != nil {
		updates["notes"] = *req.Notes
	}
	if req.Diagnosis != nil {
		updates["diagnosis"] = *req.Diagnosis
	}
	if req.TreatmentPlan != nil {
		updates["treatment_plan"] = *req.TreatmentPlan
	}
	ctx := context.Background()

	return s.appointmentRepo.Update(ctx, int32(id), updates)
}

func (s *AppointmentService) DeleteAppointment(id int) error {
	ctx := context.Background()
	return s.appointmentRepo.Delete(ctx, int32(id))
}

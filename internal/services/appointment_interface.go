package services

import (
	"github.com/prem0x01/hospital/internal/domain"
)

type IAppointmentService interface {
	GetAppointments(limit, offset int, userRole string, userID int) ([]domain.Appointment, error)
	GetAppointment(id int) (*domain.Appointment, error)
	CreateAppointment(req *domain.CreateAppointmentRequest, createdBy int) (*domain.Appointment, error)
	UpdateAppointment(id int, req *domain.UpdateAppointmentRequest) error
	DeleteAppointment(id int) error
}

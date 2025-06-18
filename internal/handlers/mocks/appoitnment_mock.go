package handlers_test

import (
	"context"
	"errors"

	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/services"
)

type MockAppointmentService struct{}

func (m *MockAppointmentService) GetAppointments(limit, offset int, userRole string, userID int) ([]domain.Appointment, error) {
	return []domain.Appointment{
		{ID: 1, Status: "Scheduled"},
	}, nil
}

func (m *MockAppointmentService) GetAppointment(id int) (*domain.Appointment, error) {
	if id == 999 {
		return nil, errors.New("not found")
	}
	return &domain.Appointment{ID: id, Status: "Scheduled"}, nil
}

func (m *MockAppointmentService) CreateAppointment(req *domain.CreateAppointmentRequest, createdBy int) (*domain.Appointment, error) {
	return &domain.Appointment{ID: 123}, nil
}

func (m *MockAppointmentService) UpdateAppointment(id int, req *domain.UpdateAppointmentRequest) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return nil
}

func (m *MockAppointmentService) DeleteAppointment(id int) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return nil
}

package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/services"
	"github.com/stretchr/testify/mock"
)

// MockAppointmentService must implement all methods of services.AppointmentService
type MockAppointmentService struct {
	mock.Mock
}

var _ services.AppointmentService = &MockAppointmentService{}

// Implement all interface methods
func (m *MockAppointmentService) GetAll(ctx context.Context, limit, offset int32) ([]domain.Appointment, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Appointment), args.Error(1)
}

func (m *MockAppointmentService) GetByID(ctx context.Context, id int32) (*domain.Appointment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Appointment), args.Error(1)
}

func (m *MockAppointmentService) Create(ctx context.Context, req *domain.CreateAppointmentRequest) (*domain.Appointment, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Appointment), args.Error(1)
}

func (m *MockAppointmentService) Update(ctx context.Context, id int32, req *domain.UpdateAppointmentRequest) error {
	args := m.Called(ctx, id, req)
	return args.Error(0)
}

func setupTestRouter() (*gin.Engine, *MockAppointmentService) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockAppointmentService)

	handler := NewAppointmentHandler(mockService)

	router := gin.New()
	router.GET("/appointments", handler.GetAppointments)
	router.POST("/appointments", handler.CreateAppointment)
	router.GET("/appointments/:id", handler.GetAppointment)
	router.PUT("/appointments/:id", handler.UpdateAppointment)

	return router, mockService
}

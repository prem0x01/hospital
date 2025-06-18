package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/handlers"
)

// Reuse this mock from earlier
type MockAppointmentService struct{}

func (m *MockAppointmentService) GetAppointments(limit, offset int, userRole string, userID int) ([]domain.Appointment, error) {
	return []domain.Appointment{{ID: 1, Status: "scheduled"}}, nil
}

func (m *MockAppointmentService) GetAppointment(id int) (*domain.Appointment, error) {
	if id == 999 {
		return nil, domain.ErrNotFound
	}
	return &domain.Appointment{ID: id, Status: "scheduled"}, nil
}

func (m *MockAppointmentService) CreateAppointment(req *domain.CreateAppointmentRequest, createdBy int) (*domain.Appointment, error) {
	return &domain.Appointment{ID: 123}, nil
}

func (m *MockAppointmentService) UpdateAppointment(id int, req *domain.UpdateAppointmentRequest) error {
	return nil
}

func (m *MockAppointmentService) DeleteAppointment(id int) error {
	return nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := &MockAppointmentService{}
	appointmentHandler := handlers.NewAppointmentHandler(mockService)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/appointments", appointmentHandler.GetAppointments)
		v1.GET("/appointments/:id", appointmentHandler.GetAppointment)
		v1.POST("/appointments", appointmentHandler.CreateAppointment)
		v1.PUT("/appointments/:id", appointmentHandler.UpdateAppointment)
		v1.DELETE("/appointments/:id", appointmentHandler.DeleteAppointment)
	}

	return router
}

func TestGetAppointments(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/v1/appointments", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetAppointment(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/v1/appointments/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateAppointment(t *testing.T) {
	router := setupRouter()

	reqBody := domain.CreateAppointmentRequest{
		PatientID:       1,
		DoctorID:        nil,
		AppointmentDate: "2025-06-17T14:00",
		Notes:           nil,
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/appointments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestUpdateAppointment(t *testing.T) {
	router := setupRouter()

	reqBody := domain.UpdateAppointmentRequest{
		Status: domain.StringPtr("Completed"),
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", "/api/v1/appointments/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteAppointment(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/v1/appointments/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

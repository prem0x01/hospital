package handlers_test

import (
	"bytes"
	"encoding/json"
	//"errors"
	"net/http"
	"net/http/httptest"
	//"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/handlers"
	//"github.com/prem0x01/hospital/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockPatientService struct {
	mock.Mock
}

func (m *MockPatientService) GetPatients(limit, offset int) ([]domain.Patient, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]domain.Patient), args.Error(1)
}

func (m *MockPatientService) GetPatient(id int) (*domain.Patient, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Patient), args.Error(1)
}

func (m *MockPatientService) CreatePatient(req *domain.CreatePatientRequest, userID int) (*domain.Patient, error) {
	args := m.Called(req, userID)
	return args.Get(0).(*domain.Patient), args.Error(1)
}

func (m *MockPatientService) UpdatePatient(id int, req *domain.UpdatePatientRequest) error {
	args := m.Called(id, req)
	return args.Error(0)
}

func (m *MockPatientService) DeletePatient(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetPatients(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.GET("/patients", handler.GetPatients)

	mockPatients := []domain.Patient{{ID: 1, Name: "John Doe"}}
	mockService.On("GetPatients", 10, 0).Return(mockPatients, nil)

	req, _ := http.NewRequest("GET", "/patients?limit=10&offset=0", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestGetPatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.GET("/patients/:id", handler.GetPatient)

	mockPatient := &domain.Patient{ID: 1, Name: "Alice"}
	mockService.On("GetPatient", 1).Return(mockPatient, nil)

	req, _ := http.NewRequest("GET", "/patients/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCreatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.POST("/patients", func(c *gin.Context) {
		c.Set("user_id", 1)
		handler.CreatePatient(c)
	})

	patientReq := &domain.CreatePatientRequest{Name: "Bob", Age: 30}
	patientRes := &domain.Patient{ID: 2, Name: "Bob", Age: 30}

	mockService.On("CreatePatient", patientReq, 1).Return(patientRes, nil)

	body, _ := json.Marshal(patientReq)
	req, _ := http.NewRequest("POST", "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockService.AssertExpectations(t)
}

func TestUpdatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.PUT("/patients/:id", handler.UpdatePatient)

	updateReq := &domain.UpdatePatientRequest{Name: "Updated Bob"}
	mockService.On("UpdatePatient", 1, updateReq).Return(nil)

	body, _ := json.Marshal(updateReq)
	req, _ := http.NewRequest("PUT", "/patients/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestDeletePatient_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.DELETE("/patients/:id", func(c *gin.Context) {
		c.Set("user_role", "doctor") // Not receptionist
		handler.DeletePatient(c)
	})

	req, _ := http.NewRequest("DELETE", "/patients/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
}

func TestDeletePatient_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPatientService)
	handler := handlers.NewPatientHandler(mockService)

	router := gin.Default()
	router.DELETE("/patients/:id", func(c *gin.Context) {
		c.Set("user_role", "receptionist")
		handler.DeletePatient(c)
	})

	mockService.On("DeletePatient", 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/patients/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

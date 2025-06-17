package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/services"
	"github.com/prem0x01/hospital/internal/utils"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler(patientService *services.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}

func (h *PatientHandler) GetPatients(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	patients, err := h.patientService.GetPatients(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get patients", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Patients retrieved successfully", patients))
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid patient ID", err.Error()))
		return
	}

	patient, err := h.patientService.GetPatient(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Patient not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Patient retrieved successfully", patient))
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var req domain.CreatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request", err.Error()))
		return
	}

	userID := c.GetInt("user_id")
	patient, err := h.patientService.CreatePatient(&req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create patient", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Patient created successfully", patient))
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid patient ID", err.Error()))
		return
	}

	var req domain.UpdatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request", err.Error()))
		return
	}

	if err := h.patientService.UpdatePatient(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update patient", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Patient updated successfully", nil))
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid patient ID", err.Error()))
		return
	}

	// Check user role - only receptionists can delete patients
	userRole := c.GetString("user_role")
	if userRole != "receptionist" {
		c.JSON(http.StatusForbidden, utils.ErrorResponse("Access denied", "Only receptionists can delete patients"))
		return
	}

	if err := h.patientService.DeletePatient(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete patient", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Patient deleted successfully", nil))
}

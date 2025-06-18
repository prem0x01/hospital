package handlers

import (
	//"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/repository"
	"github.com/prem0x01/hospital/internal/services"
	"github.com/prem0x01/hospital/internal/utils"
)

type AppointmentHandler struct {
	appointmentService services.IAppointmentService
}

func NewAppointmentHandler(appointmentService services.IAppointmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService: appointmentService}
}

func (h *AppointmentHandler) GetAppointments(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	userRole := c.GetString("user_role")
	userID := c.GetInt("user_id")

	appointments, err := h.appointmentService.GetAppointments(limit, offset, userRole, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get appointments", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Appointments retrieved successfully", appointments))
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid appointment ID", err.Error()))
		return
	}

	appointment, err := h.appointmentService.GetAppointment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Appointment not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Appointment retrieved successfully", appointment))
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var req domain.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request", err.Error()))
		return
	}

	userID := c.GetInt("user_id")
	appointment, err := h.appointmentService.CreateAppointment(&req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create appointment", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Appointment created successfully", appointment))
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid appointment ID", err.Error()))
		return
	}

	var req domain.UpdateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request", err.Error()))
		return
	}

	if err := h.appointmentService.UpdateAppointment(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update appointment", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Appointment updated successfully", nil))
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid appointment ID", err.Error()))
		return
	}

	// check user role - only receptionists can delete appointments
	userRole := c.GetString("user_role")
	if userRole != "receptionist" {
		c.JSON(http.StatusForbidden, utils.ErrorResponse("Access denied", "Only receptionists can delete appointments"))
		return
	}

	if err := h.appointmentService.DeleteAppointment(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete appointment", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Appointment deleted successfully", nil))
}

func GetDashboardStats(patientRepo *repository.PatientRepository, appointmentRepo *repository.AppointmentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		patientCount, err := patientRepo.Count(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get patient count", err.Error()))
			return
		}

		appointmentCount, err := appointmentRepo.Count(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get appointment count", err.Error()))
			return
		}

		stats := map[string]interface{}{
			"total_patients":     patientCount,
			"total_appointments": appointmentCount,
		}

		c.JSON(http.StatusOK, utils.SuccessResponse("Dashboard stats retrieved successfully", stats))
	}
}

package domain

import "time"

type Appointment struct {
	ID              int       `json:"id" db:"id"`
	PatientID       int       `json:"patient_id" db:"patient_id"`
	DoctorID        *int      `json:"doctor_id" db:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date" db:"appointment_date"`
	Status          string    `json:"status" db:"status"`
	Notes           *string   `json:"notes" db:"notes"`
	Diagnosis       *string   `json:"diagnosis" db:"diagnosis"`
	TreatmentPlan   *string   `json:"treatment_plan" db:"treatment_plan"`
	CreatedBy       *int      `json:"created_by" db:"created_by"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	PatientName     string    `json:"patient_name,omitempty"`
	DoctorName      string    `json:"doctor_name,omitempty"`
}

type CreateAppointmentRequest struct {
	PatientID       int     `json:"patient_id" binding:"required"`
	DoctorID        *int    `json:"doctor_id"`
	AppointmentDate string  `json:"appointment_date" binding:"required"`
	Notes           *string `json:"notes"`
}

type UpdateAppointmentRequest struct {
	DoctorID        *int    `json:"doctor_id"`
	AppointmentDate *string `json:"appointment_date"`
	Status          *string `json:"status"`
	Notes           *string `json:"notes"`
	Diagnosis       *string `json:"diagnosis"`
	TreatmentPlan   *string `json:"treatment_plan"`
}

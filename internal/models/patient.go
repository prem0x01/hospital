package domain

import (
	"database/sql/driver"
	"time"
)

type Patient struct {
	ID                     int       `json:"id" db:"id"`
	FirstName              string    `json:"first_name" db:"first_name"`
	LastName               string    `json:"last_name" db:"last_name"`
	Email                  *string   `json:"email" db:"email"`
	Phone                  *string   `json:"phone" db:"phone"`
	DateOfBirth            *NullDate `json:"date_of_birth" db:"date_of_birth"`
	Gender                 *string   `json:"gender" db:"gender"`
	Address                *string   `json:"address" db:"address"`
	MedicalHistory         *string   `json:"medical_history" db:"medical_history"`
	Allergies              *string   `json:"allergies" db:"allergies"`
	EmergencyContactName   *string   `json:"emergency_contact_name" db:"emergency_contact_name"`
	EmergencyContactPhone  *string   `json:"emergency_contact_phone" db:"emergency_contact_phone"`
	CreatedBy              *int      `json:"created_by" db:"created_by"`
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
}

type NullDate struct {
	Time  time.Time
	Valid bool
}

func (nd *NullDate) Scan(value interface{}) error {
	if value == nil {
		nd.Time, nd.Valid = time.Time{}, false
		return nil
	}
	nd.Valid = true
	nd.Time = value.(time.Time)
	return nil
}

func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Time, nil
}

type CreatePatientRequest struct {
	FirstName              string  `json:"first_name" binding:"required"`
	LastName               string  `json:"last_name" binding:"required"`
	Email                  *string `json:"email"`
	Phone                  *string `json:"phone"`
	DateOfBirth            *string `json:"date_of_birth"`
	Gender                 *string `json:"gender"`
	Address                *string `json:"address"`
	MedicalHistory         *string `json:"medical_history"`
	Allergies              *string `json:"allergies"`
	EmergencyContactName   *string `json:"emergency_contact_name"`
	EmergencyContactPhone  *string `json:"emergency_contact_phone"`
}

type UpdatePatientRequest struct {
	FirstName              *string `json:"first_name"`
	LastName               *string `json:"last_name"`
	Email                  *string `json:"email"`
	Phone                  *string `json:"phone"`
	DateOfBirth            *string `json:"date_of_birth"`
	Gender                 *string `json:"gender"`
	Address                *string `json:"address"`
	MedicalHistory         *string `json:"medical_history"`
	Allergies              *string `json:"allergies"`
	EmergencyContactName   *string `json:"emergency_contact_name"`
	EmergencyContactPhone  *string `json:"emergency_contact_phone"`
}


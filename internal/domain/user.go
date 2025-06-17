package domain

import (
	//	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           int32            `json:"id" db:"id"`
	Email        string           `json:"email" db:"email"`
	PasswordHash string           `json:"-" db:"password_hash"`
	Role         string           `json:"role" db:"role"`
	FirstName    string           `json:"first_name" db:"first_name"`
	LastName     string           `json:"last_name" db:"last_name"`
	CreatedAt    pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role" binding:"required,oneof=receptionist doctor"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

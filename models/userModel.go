package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserDB struct {
	Id            pgtype.Numeric `json:"id" db:"id"`
	FirstName     pgtype.Text    `json:"first_name" db:"first_name"`
	LastName      pgtype.Text    `json:"last_name" db:"last_name"`
	Email         pgtype.Text    `json:"email" db:"email"`
	ContactNumber pgtype.Text    `json:"contact_number" db:"contact_number"`
	DateOfBirth   time.Time      `json:"date_of_birth" db:"date_of_birth"`
	Role          pgtype.Numeric `json:"role" db:"role"`
	RoleName      pgtype.Text    `json:"role_name" db:"role_name"`
	IsActive      pgtype.Bool    `json:"is_active" db:"is_active"`
	Hash          pgtype.Text    `json:"hash" db:"hash"`
	Salt          pgtype.Text    `json:"salt" db:"salt"`
}

type UserSignUpDTO struct {
	Id            int    `json:"id"`
	FirstName     string `json:"first_name" validate:"required,min=2,max=20"`
	LastName      string `json:"last_name" validate:"required,min=2,max=20"`
	Email         string `json:"email" validate:"required,email"`
	ContactNumber string `json:"contact_number" validate:"len=10"`
	DateOfBirth   string `json:"date_of_birth" validate:"datetime=2006-01-02"`
	Password      string `json:"password" validate:"required"`
	Role          byte   `json:"role"`
	RoleName      string `json:"role_name"`
	IsActive      bool   `json:"is_active"`
}

type UserSignInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserDTO struct {
	FirstName     string `json:"first_name" validate:"omitempty,min=2,max=20"`
	LastName      string `json:"last_name" validate:"omitempty,min=2,max=20"`
	Email         string `json:"email" validate:"omitempty,email"`
	ContactNumber string `json:"contact_number" validate:"omitempty,len=10"`
	DateOfBirth   string `json:"date_of_birth" validate:"omitempty,datetime=2006-01-02"`
	Role          byte   `json:"role"`
	IsActive      bool   `json:"is_active"`
}

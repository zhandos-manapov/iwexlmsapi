package models

type User struct {
	Id            int    `json:"id" db:"id"`
	FirstName     string `json:"first_name" db:"first_name" validate:"required,min=2,max=20"`
	LastName      string `json:"last_name" db:"last_name" validate:"required,min=2,max=20"`
	Email         string `json:"email" db:"email" validate:"required,email"`
	ContactNumber string `json:"contact_number" db:"contact_number" validate:"len=10"`
	DateOfBirth   string `json:"date_of_birth" db:"date_of_birth" validate:"datetime=2006-01-02"`
	Password      string `json:"password" validate:"required"`
	Role          byte   `json:"role" db:"role"`
	RoleName      string `json:"role_name" db:"role_name"`
	IsActive      bool   `json:"is_active" db:"is_active"`
	Hash          string
	Salt          string
}

type UserLog struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdate struct {
	FirstName     string `json:"first_name" validate:"min=2,max=20"`
	LastName      string `json:"last_name" validate:"min=2,max=20"`
	Email         string `json:"email" validate:"email"`
	ContactNumber string `json:"contact_number" validate:"len=10"`
	DateOfBirth   string `json:"date_of_birth" validate:"datetime=2006-01-02"`
	Role          byte   `json:"role"`
	RoleName      string `json:"role_name"`
	IsActive      bool   `json:"is_active"`
}

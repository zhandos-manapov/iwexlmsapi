package models

import "time"

type EnrolledStudent struct {
	CycleID            int       `json:"cycle_id" db:"cycle_id"`
	StudentID          int       `json:"student_id" db:"student_id"`
	EnrollmentDate     time.Time `json:"enrollment_date" db:"enrollment_date"`
	Cancelled          bool      `json:"cancelled" db:"cancelled"`
	CancellationReason string    `json:"cancellation_reason" db:"cancellation_reason"`
	FirstName          string    `json:"first_name" db:"first_name"`
	LastName           string    `json:"last_name" db:"last_name"`
	Email              string    `json:"email" db:"email"`
	ContactNumber      string    `json:"contact_number" db:"contact_number"`
	DateOfBirth        time.Time `json:"date_of_birth" db:"date_of_birth"`
	Role               byte      `json:"role" db:"role"`
	RoleName           string    `json:"role_name" db:"role_name"`
	IsActive           bool      `json:"is_active" db:"is_active"`
}

type EnrollStudentsDTO struct {
	Students []int `json:"students"`
}

type ClassDB struct {
	ID                int       `json:"id" db:"id"`
	Description       string    `json:"description" db:"description"`
	StartDate         time.Time `json:"start_date" db:"start_date"`
	EndDate           time.Time `json:"end_date" db:"end_date"`
	OpenForEnrollment bool      `json:"open_for_enrollment" db:"open_for_enrollment"`
	CourseCode        string    `json:"course_code" db:"course_code"`
	BranchName        string    `json:"branch_name" db:"branch_name"`
	CourseName        string    `json:"course_name" db:"course_name"`
	BranchID          int       `json:"branch_id" db:"branch_id"`
	CourseID          int       `json:"course_id" db:"course_id"`
	Students          int       `json:"students" db:"students"`
}

type CreateClassDTO struct {
	ID                int    `json:"id"`
	Description       string `json:"description"`
	StartDate         string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate           string `json:"end_date" validate:"required,datetime=2006-01-02"`
	OpenForEnrollment bool   `json:"open_for_enrollment"`
	CourseCode        string `json:"course_code" validate:"required"`
	BranchID          int    `json:"branch_id" validate:"required"`
	CourseID          int    `json:"course_id" validate:"required"`
}

type UpdateClassDTO struct {
	Description       string `json:"description"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	OpenForEnrollment bool   `json:"open_for_enrollment"`
	CourseCode        string `json:"course_code"`
	BranchID          int    `json:"branch_id"`
	CourseID          int    `json:"course_id"`
}

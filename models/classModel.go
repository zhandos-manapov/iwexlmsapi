package models

import "time"

type Enrollment struct {
	CycleID            int       `json:"cycle_id"`
	StudentID          int       `json:"student_id"`
	EnrollmentDate     time.Time `json:"enrollment_date"`
	Cancelled          bool      `json:"cancelled"`
	CancellationReason string    `json:"cancellation_reason"`
}


type Class struct {
	ID                int    `json:"id"`
	Description       string `json:"description"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	OpenForEnrollment bool   `json:"open_for_enrollment"`
	CourseCode        string `json:"course_code"`
	BranchName        string `json:"branch_name"`
	CourseName        string `json:"course_name"`
	BranchID          int    `json:"branch_id"`
	CourseID          int    `json:"course_id"`
}

type ClassReqBodySchema struct {
	Description       string `json:"description"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	OpenForEnrollment bool   `json:"open_for_enrollment"`
	CourseCode        string `json:"course_code"`
}
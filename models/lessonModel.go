package models

import "time"

type CreateLessonDTO struct {
	ID             int    `json:"id"`
	LessonTitle    string `json:"lesson_title" validate:"required,min=2"`
	CycleId        int    `json:"cycle_id" validate:"required"`
	StartTime      string `json:"start_time" validate:"required,datetime=2006-01-02 15:04:05"`
	EndTime        string `json:"end_time" validate:"required,datetime=2006-01-02 15:04:05"`
	Description    string `json:"description"`
	RecurrenceRule string `json:"recurrence_rule"`
}

type UpdateLessonDTO struct {
	LessonTitle    string `json:"lesson_title" validate:"omitempty,min=2"`
	CycleId        int    `json:"cycle_id"`
	StartTime      string `json:"start_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	EndTime        string `json:"end_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	Description    string `json:"description"`
	RecurrenceRule string `json:"recurrence_rule"`
}

type LessonDB struct {
	ID             int       `json:"id"`
	LessonTitle    string    `json:"lesson_title" db:"lesson_title"`
	CycleId        int       `json:"cycle_id" db:"cycle_id"`
	StartTime      time.Time `json:"start_time" db:"start_time"`
	EndTime        time.Time `json:"end_time" db:"end_time"`
	Description    string    `json:"description" db:"description"`
	RecurrenceRule string    `json:"recurrence_rule" db:"recurrence_rule"`
	CourseCode     string    `json:"course_code" db:"course_code"`
}

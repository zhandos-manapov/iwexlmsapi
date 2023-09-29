package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateLessonDTO struct {
	ID          int    `json:"id"`
	LessonTitle string `json:"lesson_title" validate:"required,min=2"`
	CycleId     int    `json:"cycle_id" validate:"required"`
	StartTime   string `json:"start_time" validate:"required,datetime=2006-01-02T15:04:05Z"`
	EndTime     string `json:"end_time" validate:"required,datetime=2006-01-02T15:04:05Z"`
	Description string `json:"description"`
}

type GetIdLesson struct {
	Id int `json:"id" db:"id"`
}

type UpdateLessonDTO struct {
	LessonTitle string `json:"lesson_title" validate:"omitempty,min=2"`
	CycleId     int    `json:"cycle_id"`
	StartTime   string `json:"start_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	EndTime     string `json:"end_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	Description string `json:"description"`
}

type LessonDB struct {
	ID          pgtype.Numeric `json:"id"`
	LessonTitle pgtype.Text    `json:"lesson_title" db:"lesson_title"`
	CycleId     pgtype.Numeric `json:"cycle_id" db:"cycle_id"`
	StartTime   time.Time      `json:"start_time" db:"start_time"`
	EndTime     time.Time      `json:"end_time" db:"end_time"`
	Description pgtype.Text    `json:"description" db:"description"`
	CourseCode  pgtype.Text    `json:"course_code" db:"course_code"`
}

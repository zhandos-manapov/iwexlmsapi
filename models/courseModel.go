package models

import "github.com/jackc/pgx/v5/pgtype"

type CourseDB struct {
	CourseId    pgtype.Text `json:"course_id" db:"course_id"`
	Name        pgtype.Text `json:"name" db:"name"`
	Level       pgtype.Text `json:"level" db:"level"`
	Description pgtype.Text `json:"description" db:"description"`
	Agenda      pgtype.Text `json:"agenda" db:"agenda"`
	LevelName   pgtype.Text `json:"level_name" db:"level_name"`
}

type CreateCourseDTO struct {
	ID          int    `json:"id"`
	Agenda      string `json:"agenda"`
	Level       string `json:"level" validate:"required"`
	Description string `json:"description"`
	Name        string `json:"name" validate:"required,min=2"`
}

type UpdateCourseDTO struct {
	Agenda      string `json:"agenda"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

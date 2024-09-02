package models

import "github.com/jackc/pgx/v5/pgtype"

type LevelDB struct {
	ID        pgtype.Numeric `json:"id" db:"id"`
	LevelName pgtype.Text    `json:"level_name" db:"level_name"`
}

type CreateLevelDTO struct {
	ID        int    `json:"id"`
	LevelName string `json:"level_name" validate:"required,min=2,max=20"`
}

type UpdateLevelDTO struct {
	LevelName string `json:"level_name" validate:"omitempty,min=2,max=20"`
}

package models

type Level struct {
	ID        int    `json:"id" db:"id"`
	LevelName string `json:"level_name" db:"level_name" validate:"required,min=2,max=20"`
}

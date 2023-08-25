package models

type Level struct {
	ID        int    `json:"id"`
	LevelName string `json:"level_name" validate:"required,min=2,max=20"`
}

package models

type LevelDB struct {
	ID        int    `json:"id" db:"id"`
	LevelName string `json:"level_name" db:"level_name"`
}

type CreateLevelDTO struct {
	ID        int    `json:"id"`
	LevelName string `json:"level_name" validate:"required,min=2,max=20"`
}

type UpdateLevelDTO struct {
	LevelName string `json:"level_name" validate:"omitempty,min=2,max=20"`
}

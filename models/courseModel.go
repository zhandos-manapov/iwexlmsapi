package models

type Course struct {
	CourseId    string `json:"course_id" db:"course_id"`
	Name        string `json:"name" db:"name"`
	Level       string `json:"level" db:"level"`
	Description string `json:"description" db:"description"`
	Agenda      string `json:"agenda" db:"agenda"`
	LevelName   string `json:"level_name" db:"level_name"`
}

type CourseCreate struct {
	Agenda      string `json:"agenda"`
	Level       string `json:"level" validate:"required"`
	Description string `json:"description"`
	Name        string `json:"name" validate:"required,min=2"`
}

type CourseUpdate struct {
	Agenda      string `json:"agenda"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

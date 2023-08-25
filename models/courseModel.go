package models

type Course struct {
	CourseId    string `json:"course_id"`
	Name        string `json:"name"`
	Level       string `json:"level"`
	Description string `json:"description"`
	Agenda      string `json:"agenda"`
	LevelName   string `json:"level_name"`
}

type CourseSend struct {
	Agenda      string `json:"agenda"`
	Level       int `json:"level"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

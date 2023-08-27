package models

import "time"

type Lesson struct {
	LessonTitle string      `json:"lesson_title"`
	StartTime   time.Time   `json:"start_time" validate:"datetime=2006-01-02"`
	EndTime     time.Time   `json:"end_time" validate:"datetime=2006-01-02"`
	Description interface{} `json:"description"`
	CourseCode  string      `json:"course_code"`
}

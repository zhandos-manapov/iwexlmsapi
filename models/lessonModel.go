package models

import "time"

type Lesson struct {
	LessonTitle    string      `json:"lesson_title" db:"lesson_title"`
	StartTime      time.Time   `json:"start_time" db:"start_time"`
	EndTime        time.Time   `json:"end_time" db:"end_time"`
	Description    interface{} `json:"description" db:"description"`
	RecurrenceRule interface{} `json:"recurrence_rule" db:"recurrence_rule"`
	CourseCode     string      `json:"course_code" db:"course_code"`
}

type CreateLesson struct {
	LessonTitle    string      `json:"lesson_title" db:"lesson_title" validate:"min=2"`
	CycleId        string      `json:"cycle_id" db:"cycle_id"`
	StartTime      string      `json:"start_time" db:"start_time" validate:"datetime=2006-01-02 15:04:05"`
	EndTime        string      `json:"end_time" db:"end_time" validate:"datetime=2006-01-02 15:04:05"`
	Description    interface{} `json:"description" db:"description"`
	RecurrenceRule interface{} `json:"recurrence_rule" db:"recurrence_rule"`
}

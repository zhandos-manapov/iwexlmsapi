package models

import "time"

type Attendance struct {
	LessonId  int  `json:"lesson_id" db:"lesson_id"`
	StudentId int  `json:"student_id" db:"student_id"`
	Attended  bool `json:"attended" db:"attended"`
}

type AttendanceFind struct {
	StartTime   time.Time `json:"start_time" db:"start_time"`
	LessonTitle string    `json:"lesson_title" db:"lesson_title"`
	Id          int       `json:"id" db:"id"`
	Attended    bool      `json:"attended" db:"attended"`
}

type UpdAttendance struct {
	LessonId  int  `json:"lesson_id" db:"lesson_id"`
	StudentId int  `json:"student_id" db:"student_id"`
	Attended  bool `json:"attended" db:"attended"`
}

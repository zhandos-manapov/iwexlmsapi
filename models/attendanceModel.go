package models

type Attendance struct {
	LessonId  int    `json:"lesson_id" db:"lesson_id"`
	StudentId int    `json:"student_id" db:"student_id"`
	Attended  string `json:"attended" db:"attended"`
	Id        int    `json:"id" db:"id"`
}

type UpdAttendance struct {
	LessonId  int    `json:"lesson_id" db:"lesson_id"`
	StudentId int    `json:"student_id" db:"student_id"`
	Attended  string `json:"attended" db:"attended"`
}

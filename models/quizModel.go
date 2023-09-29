package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type QuizDB struct {
	ID                         pgtype.Numeric     `json:"id" db:"id"`
	CycleId                    pgtype.Text        `json:"cycle_id" db:"cycle_id"`
	QuizName                   pgtype.Text        `json:"quiz_name" db:"quiz_name"`
	Graded                     pgtype.Bool        `json:"graded" db:"graded"`
	Instructions               pgtype.Text        `json:"instructions" db:"instructions"`
	ShuffleAnswers             pgtype.Bool        `json:"shuffle_answers" db:"shuffle_answers"`
	IsTimed                    pgtype.Bool        `json:"is_timed" db:"is_timed"`
	TimeLimit                  pgtype.Interval    `json:"time_limit" db:"time_limit"`
	MultipleAttempts           pgtype.Bool        `json:"multiple_attempts" db:"multiple_attempts"`
	ShowResultsToStudents      pgtype.Bool        `json:"show_results_to_students" db:"show_results_to_students"`
	ShowOnlyOnce               pgtype.Bool        `json:"show_only_once" db:"show_only_once"`
	ShowCorrectAnswer          pgtype.Bool        `json:"show_correct_answer" db:"show_correct_answer"`
	ShowCorrectAnswerAt        pgtype.Timestamptz `json:"show_correct_answer_at" db:"show_correct_answer_at"`
	HideCorrectAnswerAt        pgtype.Timestamptz `json:"hide_correct_answer_at" db:"hide_correct_answer_at"`
	OneQuestionAtTime          pgtype.Bool        `json:"one_question_at_time" db:"one_question_at_time"`
	LockQuestionAfterAnswering pgtype.Bool        `json:"lock_question_after_answering" db:"lock_question_after_answering"`
	DueDate                    pgtype.Timestamptz `json:"due_date" db:"due_date"`
	AvailableFrom              pgtype.Timestamptz `json:"available_from" db:"available_from"`
	AvailableUntil             pgtype.Timestamptz `json:"available_until" db:"available_until"`
	Published                  pgtype.Bool        `json:"published" db:"published"`
	PublishedAt                pgtype.Timestamptz `json:"published_at" db:"published_at"`
}

type CreateQuizDTO struct {
	ID       int    `json:"id"`
	QuizName string `json:"quiz_name" validate:"required,min=2,max=50"`
}

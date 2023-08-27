package lesson

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT lesson.lesson_title, lesson.start_time, lesson.end_time, lesson.description, course_cycle.course_code  FROM lesson
  INNER JOIN course_cycle ON lesson.cycle_id = course_cycle.id
  WHERE lesson.id = $1`
	var lesson models.Lesson

	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&lesson.LessonTitle, &lesson.StartTime, &lesson.EndTime, &lesson.Description, &lesson.CourseCode); err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(fiber.StatusBadRequest, "Урок не найден")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Ошибка сервера: "+err.Error())
	}
	return c.JSON(lesson)
}

func FindMany(c *fiber.Ctx) error {
	query := `SELECT lesson.lesson_title, lesson.start_time, lesson.end_time, lesson.description, course_cycle.course_code  FROM lesson
  INNER JOIN course_cycle ON lesson.cycle_id = course_cycle.id`
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return err
	}

	lessons := make([]models.Lesson, 0)

	for rows.Next() {
		var lesson models.Lesson
		if err := rows.Scan(&lesson.LessonTitle, &lesson.StartTime, &lesson.EndTime, &lesson.Description); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Ошибка сервер: "+err.Error())
		}
		lessons = append(lessons, lesson)
	}

	if len(lessons) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Что то пошло не так"+err.Error())
	}
	return c.JSON(lessons)
}

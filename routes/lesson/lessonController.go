package lesson

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func getIdLesson(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT lesson.id
	FROM lesson
		INNER JOIN course_cycle ON lesson.cycle_id = course_cycle.id
	WHERE course_cycle.id = $1`
	lesson := models.GetIdLesson{}

	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&lesson.Id,
	); err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(fiber.StatusBadRequest, "Такого класса нет")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Ошибка сервера: "+err.Error())
	}
	return c.JSON(lesson)
}

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT lesson.id,
		lesson.lesson_title,
		lesson.cycle_id,
		lesson.start_time,
		lesson.end_time,
		COALESCE(lesson.description, '') as description,
		course_cycle.course_code
	FROM lesson
		INNER JOIN course_cycle ON lesson.cycle_id = course_cycle.id
	WHERE lesson.id = $1`
	lesson := models.LessonDB{}

	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&lesson.ID,
		&lesson.LessonTitle,
		&lesson.CycleId,
		&lesson.StartTime,
		&lesson.EndTime,
		&lesson.Description,
		&lesson.CourseCode,
	); err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(fiber.StatusBadRequest, "Урок не найден")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Ошибка сервера: "+err.Error())
	}
	return c.JSON(lesson)
}

func findMany(c *fiber.Ctx) error {
	query := `
	SELECT lesson.id
		lesson.lesson_title,
		lesson.cycle_id,
		lesson.start_time,
		lesson.end_time,
		COALESCE(lesson.description, '') as description,
		course_cycle.course_code
	FROM lesson
		INNER JOIN course_cycle ON lesson.cycle_id = course_cycle.id`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	lessons, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.LessonDB])
	if err != nil {
		return err
	}
	return c.JSON(lessons)
}

func createOne(c *fiber.Ctx) error {
	lesson := c.Locals("body").(*models.CreateLessonDTO)
	query := `
	INSERT INTO lesson(
    cycle_id,
    lesson_title,
    start_time,
    end_time,
    description
  )
	VALUES($1, $2, $3, $4, $5)
	RETURNING id`
	if err := database.Pool.QueryRow(
		context.Background(),
		query,
		lesson.CycleId,
		lesson.LessonTitle,
		lesson.StartTime,
		lesson.EndTime,
		lesson.Description,
		lesson.RecurrenceRule,
	).Scan(&lesson.ID); err != nil {
		return err
	}
	return c.JSON(lesson)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	lesson := c.Locals("body").(*models.UpdateLessonDTO)

	if (*lesson == models.UpdateLessonDTO{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления урока")
	}

	query := strings.Builder{}
	query.WriteString("UPDATE lesson SET")
	queryParams := []any{id}

	if lesson.CycleId != 0 {
		query.WriteString(fmt.Sprintf(" cycle_id=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, lesson.CycleId)
	}

	if lesson.LessonTitle != "" {
		query.WriteString(fmt.Sprintf(" lesson_title=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, lesson.LessonTitle)
	}

	if lesson.StartTime != "" {
		query.WriteString(fmt.Sprintf(" start_time=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, lesson.StartTime)
	}

	if lesson.EndTime != "" {
		query.WriteString(fmt.Sprintf(" end_time=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, lesson.EndTime)
	}

	if lesson.Description != "" {
		query.WriteString(fmt.Sprintf(" description=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, lesson.Description)
	}
	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE id=$1"

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Урок не найден")
	}
	return c.JSON(models.RespMsg{Message: "Урок успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM lesson WHERE id = $1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Урок не найден")
	}
	return c.JSON(models.RespMsg{Message: "Урок успешно удален"})
}

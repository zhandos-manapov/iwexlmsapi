package course

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"strings"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT course.name,
		course.course_id,
		course.level,
		course.description,
		course.agenda,
		level.level_name
	FROM course
		INNER JOIN level ON course.level = level.id
	WHERE course_id = $1`
	course := models.Course{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&course.Name,
		&course.CourseId,
		&course.Level,
		&course.Description,
		&course.Agenda,
		&course.LevelName); err != nil {
		return err
	}
	return c.JSON(course)
}

func findMany(c *fiber.Ctx) error {
	query := `
	SELECT course_id,
		course.name,
		course.level,
		course.description,
		course.agenda,
		level.level_name
	FROM course
		INNER JOIN level ON course.level = level.id`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	courses, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Course])
	if err != nil {
		return err
	}
	return c.JSON(courses)
}

func createOne(c *fiber.Ctx) error {
	course := c.Locals("body").(*models.CreateCourse)
	query := `
	INSERT INTO course (name, level, description, agenda)
	VALUES ($1, $2, $3, $4)`
	if tag, err := database.Pool.Exec(context.Background(), query, course.Name, course.Level, course.Description, course.Agenda); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Курс успешно создан"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	course := c.Locals("body").(*models.UpdateCourse)

	if course.Name == "" && course.Level == 0 && course.Description == "" && course.Agenda == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}

	query := strings.Builder{}
	query.WriteString("UPDATE course SET")
	queryParams := []any{id}

	if course.Name != "" {
		query.WriteString(fmt.Sprintf(" name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, course.Name)
	}

	if course.Level != 0 {
		query.WriteString(fmt.Sprintf(" level=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, course.Level)
	}

	if course.Description != "" {
		query.WriteString(fmt.Sprintf(" description=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, course.Description)
	}

	if course.Agenda != "" {
		query.WriteString(fmt.Sprintf(" agenda=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, course.Agenda)
	}
	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE course_id=$1"

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Курс не найден")
	}
	return c.JSON(models.RespMsg{Message: "Курс успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM course WHERE course_id = $1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Курс не найден")
	}
	return c.JSON(models.RespMsg{Message: "Курс успешно удален"})
}

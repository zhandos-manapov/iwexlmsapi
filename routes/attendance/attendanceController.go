package attendance

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT * FROM attendance WHERE lesson_id = $1`
	attendance := models.Attendance{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&attendance.LessonId,
		&attendance.StudentId,
		&attendance.Attended,
	); err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(fiber.StatusBadRequest, "Информация не найдена")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Ошибка сервера: "+err.Error())
	}
	return c.JSON(attendance)
}

func findMany(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT lesson.start_time,
		attendance.attended,
		lesson.lesson_title,
		lesson.id
		FROM lesson
		INNER JOIN attendance ON lesson.id = attendance.lesson_id
		WHERE lesson.cycle_id=$1`
	rows, err := database.Pool.Query(context.Background(), query, id)
	defer rows.Close()
	if err != nil {
		return err
	}
	courses, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.AttendanceFind])
	if err != nil {
		return err
	}
	return c.JSON(courses)
}

func createMany(c *fiber.Ctx) error {
	attendance := c.Locals("body").([]models.UpdAttendance)

	query := `
	INSERT INTO attendance(
		lesson_id,
		student_id,
		attended)
	VALUES`

	args := []any{}

	for i, attended := range attendance {
		args = append(args,
			attended.LessonId,
			attended.StudentId,
			attended.Attended,
		)
		query += fmt.Sprintf("($%d, $%d, $%d),", i*3+1, i*3+2, i*3+3)
	}

	query = query[:len(query)-1]

	if tag, err := database.Pool.Exec(context.Background(), query, args...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Успешно добавлено"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	attendance := c.Locals("body").(*models.UpdAttendance)

	if attendance.Attended == false {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}

	query := `UPDATE attendance SET attended=$2 WHERE lesson_id=$1 AND student_id=$3`

	if tag, err := database.Pool.Exec(context.Background(), query, id, attendance.Attended, attendance.StudentId); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Отчёт не найден")
	}
	return c.JSON(models.RespMsg{Message: "Отчёт успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM attendance WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Отчёт не найден")
	}
	return c.JSON(models.RespMsg{Message: "Отчёт успешно удален"})
}

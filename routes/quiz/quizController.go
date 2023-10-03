package quiz

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func findManyByCycleId(c *fiber.Ctx) error {
	cycle_id := c.Query("cycle_id")
	if cycle_id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "ID класса не указан")
	}
	query := "SELECT * FROM quiz WHERE cycle_id = $1"
	rows, err := database.Pool.Query(context.Background(), query, cycle_id)
	defer rows.Close()
	if err != nil {
		return err
	}
	quizzes, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.QuizDB])
	if err != nil {
		return err
	}
	if len(quizzes) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Тесты не найдены")
	}
	return c.JSON(quizzes)
}

func createOne(c *fiber.Ctx) error {
	quiz := c.Locals("body").(*models.CreateQuizDTO)
	query := "INSERT INTO quiz (cycle_id, quiz_name) VALUES($1, $2) RETURNING id"
	if err := database.Pool.QueryRow(context.Background(), query, quiz.CycleId, quiz.QuizName).Scan(&quiz.ID); err != nil {
		return err
	}
	return c.JSON(quiz)
}

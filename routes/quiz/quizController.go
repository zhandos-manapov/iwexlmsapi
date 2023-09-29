package quiz

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func findMany(c *fiber.Ctx) error {
	query := `
	SELECT * FROM quiz`
	rows, err := database.Pool.Query(context.Background(), query)
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
	query := "INSERT INTO quiz (quiz_name) VALUES($1) RETURNING id"
	if err := database.Pool.QueryRow(context.Background(), query, quiz.QuizName).Scan(&quiz.ID); err != nil {
		return err
	}
	return c.JSON(quiz)
}

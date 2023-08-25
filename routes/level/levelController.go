package level

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM level WHERE id=$1"
	row := database.Pool.QueryRow(context.Background(), query, id)

	var level models.Level
	if err := row.Scan(&level.ID, &level.LevelName); err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Уровень не найден",
			})
		}
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(http.StatusOK).JSON(level)
}

func FindMany(c *fiber.Ctx) error {
	query := "SELECT * FROM level"
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer rows.Close()

	levels := make([]models.Level, 0)
	for rows.Next() {
		var level models.Level
		if err := rows.Scan(&level.ID, &level.LevelName); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
		levels = append(levels, level)
	}

	if len(levels) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Что-то пошло не так",
		})
	}

	return c.Status(http.StatusOK).JSON(levels)
}

func CreateOne(c *fiber.Ctx) error {
	var level models.Level
	if err := c.BodyParser(&level); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "INSERT INTO level (level_name) VALUES($1)"
	_, err := database.Pool.Exec(context.Background(), query, level.LevelName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Такое название уровня уже существует",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Создано успешно"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var level models.Level
	if err := c.BodyParser(&level); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "UPDATE level SET level_name=$1 WHERE id=$2"
	_, err := database.Pool.Exec(context.Background(), query, level.LevelName, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка изменения данных",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Успешно обновлено"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM level WHERE id=$1"
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Такого уровня не существует",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Успешно удалено"})
}

package users

import (
	"context"
	"iwexlmsapi/database"

	"github.com/gofiber/fiber/v2"
)

func Toggle(c *fiber.Ctx) error {
	id := c.Params("id")
	var requestBody struct {
		IsActive bool `json:"is_active"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing data",
		})
	}

	query := `
        UPDATE users SET is_active=$1 WHERE id=$2`
	_, err := database.Pool.Exec(context.Background(), query, requestBody.IsActive, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	message := "Успешно активирован"
	if !requestBody.IsActive {
		message = "Успешно деактивирован"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

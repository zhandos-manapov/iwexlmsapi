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
		err := fiber.NewError(fiber.StatusBadRequest, "Error parsing data")
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.Message,
		})
	}

	query := `
        UPDATE users SET is_active=$1 WHERE id=$2`
	_, err := database.Pool.Exec(context.Background(), query, requestBody.IsActive, id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	message := "Успешно активирован"
	if !requestBody.IsActive {
		message = "Успешно деактивирован"
	}

	return c.JSON(fiber.Map{"message": message})
}

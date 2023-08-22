package main

import (
	"errors"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"iwexlmsapi/xvalidator"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
}

func main() {

	loadEnv()

	database.ConnectToDB()
	defer database.DisconnectFromDB()

	validate := validator.New()
	xvalidator.InitValidator(validate)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(models.ServerError{Message: err.Error()})
		},
	})

	setupRoutes(app)
	app.Listen(":3030")
}

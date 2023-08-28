package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"iwexlmsapi/routes/files"
	"iwexlmsapi/utils"
	"iwexlmsapi/xvalidator"
	"log"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
}

func main() {

	loadEnv()

	utils.InitKeys()
	files.InitConstants()
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
			return c.Status(code).JSON(models.RespMsg{Message: err.Error()})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	setupRoutes(app)
	app.Listen(":3030")
}

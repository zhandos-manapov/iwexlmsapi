package level

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupLevelRouter(router *fiber.Router) {
	levelRouter := (*router).Group("/levels")
	levelRouter.Get("/:id", FindOne)
	levelRouter.Get("/", FindMany)
	levelRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Level{}), CreateOne)
	levelRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.Level{}), UpdateOne)
	levelRouter.Delete("/:id", DeleteOne)
}

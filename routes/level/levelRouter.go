package level

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupLevelRouter(router *fiber.Router) {
	levelRouter := (*router).Group("/levels")
	levelRouter.Get("/:id", findOne)
	levelRouter.Get("/", findMany)
	levelRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateLevelDTO{}), createOne)
	levelRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateLevelDTO{}), updateOne)
	levelRouter.Delete("/:id", deleteOne)
}

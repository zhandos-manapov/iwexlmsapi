package level

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupLevelRoute(router fiber.Router) {
	localRouter := router.Group("/levels")
	localRouter.Get("/:id", FindOne)
	localRouter.Get("/", FindMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Level{}), CreateOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.Level{}), UpdateOne)
	localRouter.Delete("/:id", DeleteOne)
}

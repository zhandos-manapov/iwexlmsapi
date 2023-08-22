package city

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCityRoutes(router fiber.Router) {
	cityRouter := router.Group("/cities")

	cityRouter.Get("/", FindAll)
	cityRouter.Get("/:id", FindOne)
	cityRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.City{}), CreateOne)
	cityRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.City{}), UpdateOne)
	cityRouter.Delete("/:id", DeleteOne)
}

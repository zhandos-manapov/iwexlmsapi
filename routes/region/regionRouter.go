package region

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRegionRoutes(router *fiber.Router) {
	regionRouter := (*router).Group("/regions")

	regionRouter.Get("/:id", FindOne)
	regionRouter.Get("/", FindAll)
	regionRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Region{}), CreateOne)
	regionRouter.Put("/:id", middleware.BodyParserValidatorMiddleware(&models.Region{}), UpdateOne)
	regionRouter.Delete("/:id", DeleteOne)
}

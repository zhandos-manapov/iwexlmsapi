package region

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRegionRouter(router *fiber.Router) {
	regionRouter := (*router).Group("/regions")

	regionRouter.Get("/:id", FindOne)
	regionRouter.Get("/", FindMany)
	regionRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Region{}), CreateOne)
	regionRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.Region{}), UpdateOne)
	regionRouter.Delete("/:id", DeleteOne)
}

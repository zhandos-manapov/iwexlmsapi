package region

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRegionRouter(router *fiber.Router) {
	regionRouter := (*router).Group("/regions")

	regionRouter.Get("/:id", findOne)
	regionRouter.Get("/", findMany)
	regionRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Region{}), createOne)
	regionRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.Region{}), updateOne)
	regionRouter.Delete("/:id", deleteOne)
}

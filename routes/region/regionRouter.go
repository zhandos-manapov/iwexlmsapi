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
	regionRouter.Post("/", middleware.BodyParserValidatorMiddlewareForStruct(&models.CreateRegionDTO{}), createOne)
	regionRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdateRegionDTO{}), updateOne)
	regionRouter.Delete("/:id", deleteOne)
}

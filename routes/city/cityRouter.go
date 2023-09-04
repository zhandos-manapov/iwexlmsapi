package city

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCityRouter(router *fiber.Router) {
	cityRouter := (*router).Group("/cities")

	cityRouter.Get("/", findMany)
	cityRouter.Get("/:id", findOne)
	cityRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateCityDTO{}), createOne)
	cityRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateCityDTO{}), updateOne)
	cityRouter.Delete("/:id", deleteOne)
}

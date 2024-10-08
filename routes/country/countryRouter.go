package country

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCountryRouter(router *fiber.Router) {
	countryRouter := (*router).Group("/countries")

	countryRouter.Get("/:id", findOne)
	countryRouter.Get("/", findMany)
	countryRouter.Post("/", middleware.BodyParserValidatorMiddlewareForStruct(&models.CreateCountryDTO{}), createOne)
	countryRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdateCountryDTO{}), updateOne)
	countryRouter.Delete("/:id", deleteOne)
}

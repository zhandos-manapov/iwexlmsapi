package country

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCountryRouter(router *fiber.Router) {
	countryRouter := (*router).Group("/countries")

	countryRouter.Get("/:id", FindOne)
	countryRouter.Get("/", FindMany)
	countryRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Country{}), CreateOne)
	countryRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.Country{}), UpdateOne)
	countryRouter.Delete("/:id", DeleteOne)
}

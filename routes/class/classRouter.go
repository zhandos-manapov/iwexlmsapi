package class

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupClassRoutes(router *fiber.Router) {
	classRouter := (*router).Group("/classes")
	classRouter.Get("/", findMany)
	classRouter.Get("/:id", findOne)
	classRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateClass{}), createOne)
	classRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateClass{}), updateOne)
	classRouter.Delete("/:id", deleteOne)
	classRouter.Get("/:id/people", getEnrollment)
}

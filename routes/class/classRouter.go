package class

import (
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
)

func SetupClassRoutes(router *fiber.Router) {
	classRouter := (*router).Group("/classes")
	classRouter.Get("/", findMany)
	classRouter.Get("/:id", findOne)
	classRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateClass{}), createOne)
	classRouter.Patch("/:id", updateOne)
	classRouter.Delete("/:id", deleteOne)
	classRouter.Get("/:id/people", getEnrollment)
	// classRouter.Post("/:id/people", addEnrollment)
}

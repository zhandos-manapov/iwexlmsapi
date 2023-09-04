package class

import (
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
)

func SetupClassRouter(router *fiber.Router) {
	classRouter := (*router).Group("/classes")
	classRouter.Get("/", findMany)
	classRouter.Get("/:id", findOne)
	classRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateClassDTO{}), createOne)
	classRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateClassDTO{}), updateOne)
	classRouter.Delete("/:id", deleteOne)
	classRouter.Get("/:id/people", getEnrollment)
	// classRouter.Post("/:id/people", addEnrollment)
}

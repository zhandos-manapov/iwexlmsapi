package class

import (
	"github.com/gofiber/fiber/v2"
)

func SetupClassRoutes(r *fiber.Router) {
	classRouter := (*r).Group("/class")

	classRouter.Get("/", findMany)
	classRouter.Get("/:id", findOne)
	classRouter.Patch("/:id", updateOne)
	classRouter.Delete("/:id", deleteOne)

	classRouter.Post("/:id/toggle", ToggleOpenForEnrollment)


	classRouter.Get("/:id/people", getEnrollment)
	classRouter.Post("/:id/people", addEnrollment)
}

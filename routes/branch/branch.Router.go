package branch

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBranchRoutes(r *fiber.Router) {
	branchRouter := (*r).Group("/branch")

	branchRouter.Get("/", findMany)
	branchRouter.Get("/:id", findOne)
	branchRouter.Post("/", createOne)
	branchRouter.Delete("/:id", deleteOne)
	branchRouter.Patch("/:id", updateOne)
}

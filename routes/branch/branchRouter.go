package branch

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupBranchRouter(router *fiber.Router) {
	branchRouter := (*router).Group("/branches")

	branchRouter.Get("/", findMany)
	branchRouter.Get("/:id", findOne)
	branchRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateBranchDTO{}), createOne)
	branchRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateBranchDTO{}), updateOne)
	branchRouter.Delete("/:id", deleteOne)
}

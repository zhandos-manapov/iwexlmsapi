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
	branchRouter.Post("/", middleware.BodyParserValidatorMiddlewareForStruct(&models.CreateBranchDTO{}), createOne)
	branchRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdateBranchDTO{}), updateOne)
	branchRouter.Delete("/:id", deleteOne)
}

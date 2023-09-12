package files

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupFilesRouter(router *fiber.Router) {
	filesRouter := (*router).Group("/files")
	filesRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.FileOperationsReqBody{}), fileOperations)
	filesRouter.Get("/GetImage", getImage)
}

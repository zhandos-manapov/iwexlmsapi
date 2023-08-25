package files

import (
	"iwexlmsapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupFilesRoute(router *fiber.Router) {
	filesRouter := (*router).Group("/files")
	filesRouter.Post("/", middleware.BodyParserValidatorMiddleware(&fileOperationsReqBody{}), fileOperations)
}

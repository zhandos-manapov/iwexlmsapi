package quiz

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupQuizRouter(router *fiber.Router) {
	quizRouter := (*router).Group("/quizzes")

	quizRouter.Get("/", findMany)
	quizRouter.Post("/", middleware.BodyParserValidatorMiddlewareForStruct(&models.CreateQuizDTO{}), createOne)
}

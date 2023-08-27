package lesson

import (
	"github.com/gofiber/fiber/v2"
)

func SetupLessonRouter(router fiber.Router) {
	localRouter := router.Group("/lesson")
	localRouter.Get("/:id", FindOne)
	localRouter.Get("/", FindMany)
}

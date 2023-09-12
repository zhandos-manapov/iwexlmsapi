package files

import (
	"github.com/gofiber/fiber/v2"
	"path"
)

func getImage(c *fiber.Ctx) error {
	image := c.Query("path")
	return c.SendFile(path.Join(CONTENT_ROOT_PATH, image))
}

package files

import (
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func fileOperations(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)

	switch body.Action {
	case "details":
		return getFileDetails(c)
	case "copy":
		return copyFiles(c)
	case "move":
		return moveFiles(c)
	case "create":
		return createFolder(c)
	case "delete":
		return deleteFolder(c)
	case "rename":
		return renameFolder(c)
	case "search":
		return searchItem(c)
	case "read":
		return readFolder(c)
	default:
		return nil
	}
}

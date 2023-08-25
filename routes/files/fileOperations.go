package files

import (
	"io/ioutil"
	"path"

	"github.com/gofiber/fiber/v2"
)

func readFiles(c *fiber.Ctx) error {
	return nil
}

func getFileDetails(c *fiber.Ctx) error {
	return nil
}

func copyFiles(c *fiber.Ctx) error {
	return nil
}

func moveFiles(c *fiber.Ctx) error {
	return nil
}

func createFolder(c *fiber.Ctx) error {
	return nil
}

func deleteFolder(c *fiber.Ctx) error {
	return nil
}

func renameFolder(c *fiber.Ctx) error {
	return nil
}

func searchItem(c *fiber.Ctx) error {
	return nil
}

func readFolder(c *fiber.Ctx) error {
	body := c.Locals("body").(*fileOperationsReqBody)
	files, err := ioutil.ReadDir(path.Join(CONTENT_ROOT_PATH, body.Path))
	if err != nil {
		return err
	}
	filesCnt := len(files)
	filesList := make([]string, filesCnt)
	for i, file := range files {
		filesList[i] = file.Name()
	}
	folderInfo, err := fileStat(body.Path)
	if err != nil {
		return err
	}
	filesInfo, err := readDirectories(filesList)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"cwd": folderInfo, "files": filesInfo})
}

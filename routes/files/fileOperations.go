package files

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"io/ioutil"
	"iwexlmsapi/models"
	"os"
	"path"
	"strings"
)

func readFiles(c *fiber.Ctx) error {
	return nil
}

func getFileDetails(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	isNamesAvailable := len(body.Names) > 0
	rootName := path.Base(CONTENT_ROOT_PATH)
	contentRootPath := path.Join(CONTENT_ROOT_PATH, body.Path)
	filterPath := body.Data[0].FilterPath

	if len(body.Names) == 0 && len(body.Data) != 0 {
		nameValues := []string{}
		for _, item := range body.Data {
			nameValues = append(nameValues, item.Name)
		}
		body.Names = nameValues
	}
	if len(body.Names) == 1 {
		filename := ""
		if isNamesAvailable {
			filename = body.Names[0]
		}
		if details, err := fileDetails(path.Join(contentRootPath, filename), filterPath); err != nil {
			return err
		} else {
			if !details.IsFile {
				if size, err := getFolderSize(path.Join(contentRootPath, filename)); err != nil {
					return err
				} else {
					details.Size = getSize(int64(size))
				}
			}
			if filterPath == "" {
				details.Location = path.Join(filterPath, body.Names[0])
			} else {
				details.Location = path.Join(rootName, filterPath, body.Names[0])
			}
			return c.JSON(fiber.Map{
				"details": details,
			})
		}
	} else {
		size := 0
		for _, item := range body.Data {
			currPath := path.Join(contentRootPath, item.Name)
			if details, err := os.Stat(currPath); err != nil {
				return err
			} else {
				if details.IsDir() {
					if currSize, err := getFolderSize(currPath); err != nil {
						return err
					} else {
						size += currSize
					}
				} else {
					size += int(details.Size())
				}
			}
		}
		details := models.FileStruct{
			Name:          strings.Join(body.Names, ", "),
			Size:          getSize(int64(size)),
			MultipleFiles: true,
			Location:      path.Join(rootName, filterPath),
		}
		return c.JSON(fiber.Map{
			"details": details,
		})
	}
}

func copyFiles(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	fileList := []models.FileDetails{}
	replaceFileList := []string{}

	for _, item := range body.Data {
		fromPath := path.Join(CONTENT_ROOT_PATH, item.FilterPath, item.Name)
		toPath := path.Join(CONTENT_ROOT_PATH, body.TargetPath, item.Name)

		copyName := item.Name
		if isRenameChecking, err := checkForFileUpdate(body, fromPath, toPath, &item, &copyName); err != nil {
			return err
		} else {
			if !isRenameChecking {
				toPath = path.Join(CONTENT_ROOT_PATH, body.TargetPath, copyName)
				if item.IsFile {
					originalFile, err := os.Open(fromPath)
					if err != nil {
						return err
					}
					defer originalFile.Close()

					newFile, err := os.Create(toPath)
					if err != nil {
						return err
					}
					defer newFile.Close()

					if _, err := io.Copy(newFile, originalFile); err != nil {
						return err
					}
				} else {
					if err := copyFolder(fromPath, toPath); err != nil {
						return err
					}
				}
				list := item
				list.FilterPath = body.TargetPath
				list.Name = copyName
				fileList = append(fileList, list)
			} else {
				replaceFileList = append(replaceFileList, item.Name)
			}
		}
	}

	if len(replaceFileList) == 0 {
		return c.JSON(fiber.Map{
			"files": fileList,
		})
	} else {
		return c.JSON(fiber.Map{
			"error": fiber.Map{
				"message":    "File Already Exists.",
				"code":       "400",
				"fileExists": replaceFileList,
			},
		})
	}
}

func moveFiles(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	fileList := []models.FileDetails{}
	replaceFileList := []string{}

	for _, item := range body.Data {
		fromPath := path.Join(CONTENT_ROOT_PATH, item.FilterPath, item.Name)
		toPath := path.Join(CONTENT_ROOT_PATH, body.TargetPath, item.Name)

		copyName := item.Name
		if isRenameChecking, err := checkForFileUpdate(body, fromPath, toPath, &item, &copyName); err != nil {
			return err
		} else {
			if !isRenameChecking {
				toPath = path.Join(CONTENT_ROOT_PATH, body.TargetPath, copyName)
				if item.IsFile {
					if err := os.Rename(fromPath, toPath); err != nil {
						return err
					}
				} else {
					if err := moveFolder(fromPath, toPath); err != nil {
						return err
					} else {
						if err := os.RemoveAll(fromPath); err != nil {
							return err
						}
					}
				}
				list := item
				list.Name = copyName
				list.FilterPath = body.TargetPath
				fileList = append(fileList, list)
			} else {
				replaceFileList = append(replaceFileList, item.Name)
			}
		}
	}

	if len(replaceFileList) == 0 {
		return c.JSON(fiber.Map{
			"files": fileList,
		})
	} else {
		return c.JSON(fiber.Map{
			"error": fiber.Map{
				"message":    "File Already Exists.",
				"code":       "400",
				"fileExists": replaceFileList,
			},
		})
	}
}

func createFolder(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	newDirectoryPath := path.Join(CONTENT_ROOT_PATH, body.Path, body.Name)
	if _, err := os.Stat(newDirectoryPath); os.IsExist(err) {
		return c.JSON(fiber.Map{
			"error": fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": fmt.Sprintf("A file or folder with the name %s already exists", body.Name),
			},
		})
	} else {
		if err := os.Mkdir(newDirectoryPath, os.ModePerm); err != nil {
			return err
		}
		folderInfo, err := fileManagerDirectoryContent(body, newDirectoryPath, "")
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"files": folderInfo,
		})
	}
}

func deleteFolder(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	data := []models.FileStruct{}
	for i := 0; i < len(body.Data); i++ {
		newDirectoryPath := path.Join(CONTENT_ROOT_PATH, body.Data[i].FilterPath, body.Data[i].Name)
		info, err := fileManagerDirectoryContent(body, newDirectoryPath, body.Data[i].FilterPath)
		data = append(data, *info)
		if err != nil {
			return err
		}
		if info.IsFile {
			if err := os.Remove(newDirectoryPath); err != nil {
				return err
			}
		} else {
			if err := os.RemoveAll(newDirectoryPath); err != nil {
				return err
			}
		}
	}
	return c.JSON(fiber.Map{
		"files": data,
	})
}

func renameFolder(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	oldName := body.Name
	newName := body.NewName
	oldDirectoryPath := path.Join(CONTENT_ROOT_PATH, body.Data[0].FilterPath, oldName)
	newDirectoryPath := path.Join(CONTENT_ROOT_PATH, body.Data[0].FilterPath, newName)
	filenameExists, err := checkForDuplicates(path.Join(CONTENT_ROOT_PATH, body.Data[0].FilterPath), newName, body.Data[0].IsFile)
	if err != nil {
		return err
	}
	if filenameExists {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("A file or folder with the name %s already exists.", body.NewName))
	} else {
		if err := os.Rename(oldDirectoryPath, newDirectoryPath); err != nil {
			return err
		}
		info, err := fileManagerDirectoryContent(body, newDirectoryPath, "")
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"files": info,
		})
	}
}

func searchItem(c *fiber.Ctx) error {
	return nil
}

func readFolder(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileOperationsReqBody)
	files, err := ioutil.ReadDir(path.Join(CONTENT_ROOT_PATH, body.Path))
	if err != nil {
		return err
	}
	filesCnt := len(files)
	filesList := make([]string, filesCnt)
	for i, file := range files {
		filesList[i] = file.Name()
	}
	folderInfo, err := fileManagerDirectoryContent(body, path.Join(CONTENT_ROOT_PATH, body.Path), "")
	if err != nil {
		return err
	}
	filesInfo, err := readDirectories(body, filesList)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"cwd": folderInfo, "files": filesInfo})
}

package files

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"iwexlmsapi/models"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

func getImage(c *fiber.Ctx) error {
	image := c.Query("path")
	return c.SendFile(path.Join(CONTENT_ROOT_PATH, image))
}

func updloadFiles(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileUploadReqBody)
	uploadObj := models.FileDetails{}
	if err := json.Unmarshal([]byte(body.Data), &uploadObj); err != nil {
		return err
	}

	if body.Action == "save" {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		for _, file := range form.File["uploadFiles"] {
			uploadedFile, err := file.Open()
			defer uploadedFile.Close()
			if err != nil {
				return err
			}
			targetPath := path.Join(CONTENT_ROOT_PATH, body.Path, file.Filename)
			fmt.Println(targetPath)

			newFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}

			if _, err := io.Copy(newFile, uploadedFile); err != nil {
				return err
			}
		}
	} else if body.Action == "remove" {
		// TO DO
	}
	return c.SendString("Success")
}

func downloadFiles(c *fiber.Ctx) error {
	body := c.Locals("body").(*models.FileDownloadReqBody)
	downloadObj := models.DownloadObj{}
	if err := json.Unmarshal([]byte(body.DownloadInput), &downloadObj); err != nil {
		return err
	}
	if len(downloadObj.Names) == 1 && downloadObj.Data[0].IsFile {
		return c.Download(path.Join(CONTENT_ROOT_PATH, downloadObj.Path, downloadObj.Names[0]))
	} else {
		buf := new(bytes.Buffer)
		zipWriter := zip.NewWriter(buf)

		for _, item := range downloadObj.Data {
			if item.IsFile {
				f1, err := os.Open(path.Join(CONTENT_ROOT_PATH, item.FilterPath, item.Name))
				if err != nil {
					return err
				}

				fileInfo, err := f1.Stat()
				if err != nil {
					return err
				}

				header, err := zip.FileInfoHeader(fileInfo)
				if err != nil {
					return err
				}
				header.Name = item.Name

				w1, err := zipWriter.CreateHeader(header)
				if err != nil {
					return err
				}

				if _, err := io.Copy(w1, f1); err != nil {
					return err
				}

				f1.Close()
			} else {

			}
		}

		if err := zipWriter.Close(); err != nil {
			return err
		}

		c.Set(fiber.HeaderContentType, "application/zip")
		c.Set(fiber.HeaderContentDisposition, `attachment; filename="files.zip"`)

		return c.Send(buf.Bytes())
	}
}

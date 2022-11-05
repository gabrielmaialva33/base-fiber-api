package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"strings"
)

type File struct {
	FileName   string `json:"filename"`
	FileFormat string `json:"format"`
	FileType   string `json:"type"`
	Size       int64  `json:"size"`
	Url        string `json:"url"`
}

func Store(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting form",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	files := form.File["files"]
	var links []*File

	for _, file := range files {
		var link File

		filename := strings.Split(strings.ReplaceAll(file.Filename, " ", "_"), ".")[0] + "_" + uuid.Must(uuid.NewV4()).String()
		size := file.Size
		fileType := strings.Split(file.Header["Content-Type"][0], "/")[0]
		format := strings.Split(file.Header["Content-Type"][0], "/")[1]

		if err := c.SaveFile(file, "public/uploads/"+filename+"."+format); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while getting form",
				"error":   err.Error(),
				"status":  fiber.StatusInternalServerError,
				"display": true,
			})
		}

		link.FileName = filename
		link.Size = size
		link.Url = c.BaseURL() + "/files/uploads/" + filename + "." + format
		link.FileFormat = format
		link.FileType = fileType

		links = append(links, &link)
	}

	return c.JSON(links)
}

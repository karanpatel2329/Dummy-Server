package main

import (
	"bytes"
	"encoding/base64"
	"os"
	"strconv"

	"github.com/go-pdf/fpdf"
	"github.com/gofiber/fiber/v3"
)

type UserBalance struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func generatePdfBase64(u UserBalance) (string, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, u.Name)
	pdf.Cell(80, 10, strconv.Itoa(u.Balance))

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	app.Get("/", func(c fiber.Ctx) error { return c.SendString("Hello, World!") })

	app.Post("/download", func(c fiber.Ctx) error {
		var user UserBalance

		if err := c.Bind().Body(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if user.Name == "" || user.Balance == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Name and Amount are required",
			})
		}

		pdfBase64, err := generatePdfBase64(user)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to generate PDF",
			})
		}

		return c.JSON(fiber.Map{
			"content": pdfBase64,
		})
	})
	app.Listen(":" + port)
}

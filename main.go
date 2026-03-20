package main

import (
	"os"
	"strconv"

	"github.com/go-pdf/fpdf"
	"github.com/gofiber/fiber/v3"
)

type UserBalance struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func generatePdf(u UserBalance) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, u.Name)
	pdf.Cell(80, 10, strconv.Itoa(u.Age))
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/download", func(c fiber.Ctx) error {
		var user UserBalance
		if err := c.Bind().Body(&user); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}
		if user.Name == "" || user.Age == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Name and Amount are required",
			})
		}
		generatePdf(user)

		return c.Res().Download("hello.pdf")
	})

	app.Listen(":" + port)
}

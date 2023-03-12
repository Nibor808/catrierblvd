package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

func main() {
	engine := amber.New("./views", ".amber")
	config := fiber.Config{
		AppName:            "CartierBlvd",
		EnableIPValidation: true,
		EnablePrintRoutes:  true,
		Views:              engine,
	}

	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Saturday Night At The Movies",
		}, "layouts/main")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

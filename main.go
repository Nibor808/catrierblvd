package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	engine := html.New("./views", ".html")

	config := fiber.Config{
		AppName:            "CartierBlvd",
		EnableIPValidation: true,
		EnablePrintRoutes:  true,
		UnescapePath:       true,
		ViewsLayout:        "index",
		Views:              engine,
	}

	app := fiber.New(config)

	// Static
	app.Static("/", "./public")

	// Middleware
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("layouts/main", fiber.Map{
			"Title": "Saturday Night At The Movies",
		})
	})

	log.Fatal(app.Listen(":3000"))
}

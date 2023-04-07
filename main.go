package main

import (
	"cartierblvd/views/textmaps"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// initialize text maps for layouts
	textmaps.Init()

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
		return c.Render("layouts/main", textmaps.GetMaps("MainMa"))
	})

	log.Fatal(app.Listen(":3000"))
}

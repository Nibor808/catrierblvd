package maps

import "github.com/gofiber/fiber/v2"

func GetMaps(name string) fiber.Map {
	htmlMaps := make(map[string]fiber.Map)

	htmlMaps["MainMap"] = fiber.Map{
		"Title":         "Saturday Night At The Movies",
		"MainTextIntro": "This is the main text intro.",
	}

	return htmlMaps[name]
}

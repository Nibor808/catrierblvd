package maps

import "github.com/gofiber/fiber/v2"

var htmlMaps = make(map[string]fiber.Map)

type GetErrorParams struct {
	name string
	key  string
}

func Init() {
	htmlMaps["MainMap"] = fiber.Map{
		"Title":         "Saturday Night At The Movies",
		"MainTextIntro": "This is the main text intro.",
	}
}

func getErrorValue(args GetErrorParams) fiber.Map {
	var returnValue string

	if len(args.key) > 0 {
		returnValue = "Map" + args.name + "." + args.key + " not found"
	} else {
		returnValue = "Map" + args.name + "not found"
	}

	return fiber.Map{
		"Error": returnValue,
	}
}

func GetMaps(name string) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return getErrorValue(GetErrorParams{name: name})
	}

	return found
}

func UpdateMap(name string, key string, value string) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return getErrorValue(GetErrorParams{name: name, key: key})
	}

	found[key] = value

	return found
}

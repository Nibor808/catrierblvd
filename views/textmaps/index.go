package textmaps

import "github.com/gofiber/fiber/v2"

type TextMap string

const (
	Main     TextMap = "MainMap"
	About    TextMap = "AboutMap"
	MapError TextMap = "MapError"
	KeyError TextMap = "KeyError"
)

type GetErrorParams struct {
	name TextMap
	key  string
}

var htmlMaps = make(map[TextMap]fiber.Map)

func Init() {
	htmlMaps[Main] = MainMap
}

func getErrorValue(args GetErrorParams) fiber.Map {
	var returnValue fiber.Map

	if len(args.key) > 0 {
		returnValue = fiber.Map{
			string(MapError): args.name,
			string(KeyError): args.key,
		}
	} else {
		returnValue = fiber.Map{
			string(MapError): args.name,
		}
	}

	return returnValue
}

func GetMaps(name TextMap) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return getErrorValue(GetErrorParams{name: name})
	}

	return found
}

func UpdateMap(name TextMap, key string, value string) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return getErrorValue(GetErrorParams{name: name, key: key})
	}

	found[key] = value

	return found
}

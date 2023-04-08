package textmaps

import (
	em "cartierblvd/errormaps"
	"cartierblvd/types"
	"github.com/gofiber/fiber/v2"
)

const (
	Main string = "MainMap"
)

var htmlMaps = make(map[string]fiber.Map)

func Init() {
	htmlMaps[Main] = MainMap
}

func GetMap(name string) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return em.GetErrorValue(types.GetErrorParams{Name: name})
	}

	return found
}

func UpdateMap(name string, key string, value string) fiber.Map {
	found, ok := htmlMaps[name]
	if !ok {
		return em.GetErrorValue(types.GetErrorParams{Name: name, Key: key})
	}

	found[key] = value

	return found
}

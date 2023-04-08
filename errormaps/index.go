package errormaps

import (
	"cartierblvd/types"
	"github.com/gofiber/fiber/v2"
)

const (
	MapError string = "MapError"
	KeyError string = "KeyError"
)

func GetErrorValue(args types.GetErrorParams) fiber.Map {
	var returnValue fiber.Map

	if len(args.Key) > 0 {
		returnValue = fiber.Map{
			MapError: args.Name,
			KeyError: args.Key,
		}
	} else {
		returnValue = fiber.Map{
			MapError: args.Name,
		}
	}

	return returnValue
}

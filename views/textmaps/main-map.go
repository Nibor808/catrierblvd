package textmaps

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func getMovieDate() string {
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	return fmt.Sprintln(month, day, year)
}

var MainMap = fiber.Map{
	"Title":      "Saturday Night At The Movies",
	"Date":       getMovieDate(),
	"Intro":      "This weeks movie",
	"MovieTitle": "The Conjuring",
	"ShowInfo":   "Show time 8:30pm sharp",
	"EventInfo":  "BYO beverages and seating. Spots are first come first serve.",
}

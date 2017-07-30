package main

import (
	"fmt"
	"os"

	"git.heroku.com/dodosoft-api/events"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/events/:id", func(context echo.Context) error {
		id := context.Param("id")
		event, err := events.QueryEvent(id)
		if err != nil {
			return context.String(500, "can't get event.")
		}
		return context.JSON(200, event)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "10259"
	}
	err := e.Start(fmt.Sprintf(":%s", port))
	fmt.Println(err)
}

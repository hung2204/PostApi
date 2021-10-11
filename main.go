package main

import (
	"post/subscribe"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	server.GET("/getData", subscribe.Get)

	server.Logger.Fatal(server.Start(":2204"))
}

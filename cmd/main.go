package main

import (
	"log"
	"net/http"

	config "github.com/CodeChefVIT/cookoff-backend/common/config"
	"github.com/CodeChefVIT/cookoff-backend/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	app.Use(middleware.Logger())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{appConfig.ClientOrigin},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	database.ConnectDB(&appConfig)
	database.RunMigrations(database.DB)

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := "Not found"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		app.Logger.Error(err)
		c.JSON(code, map[string]interface{}{
			"status":  "false",
			"code":    code,
			"message": message,
		})
	}

	app.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status":  "true",
			"message": "pong",
		})
	})

	app.Logger.Fatal(app.Start(":8080"))
}

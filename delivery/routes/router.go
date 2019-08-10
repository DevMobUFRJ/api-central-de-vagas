package routes

import (
	"api-central-de-vagas/resources/config"
	"api-central-de-vagas/resources/injection"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Routes() {
	var port string
	cfg := config.NewViperConfig()
	ctrl := injection.Controller

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/vagas/user/:userId", ctrl.FindUserById)

	port = os.Getenv("PORT")

	if port != "" {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":" + cfg.GetString("port")))
	}
}

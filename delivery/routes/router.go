package routes

import (
	"api-central-de-vagas/resources/injection"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Routes() {
	var port string
	ctrl := injection.Controller

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user/create", ctrl.CreateUser) // Post Request Body
	e.POST("/user/create/cv", ctrl.SendCurriculum) // Post Curriculum file with Token in the header

	port = os.Getenv("PORT")

	if port != "" {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":80"))
	}
}

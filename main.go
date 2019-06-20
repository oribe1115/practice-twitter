package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/oribe1115/practice-twitter/handler"
	"github.com/oribe1115/practice-twitter/model"
)

func main() {
	enviroment := os.Getenv("ENVIROMENT")
	if enviroment == "" {
		model.LoadEnv()
	}

	// model.GetTwitterAPI()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})

	e.GET("/authorize", handler.GetRequestTokenHandler)
	e.GET("/authorize/callback", handler.GetAccessTokenHandler)

	e.GET("/search", handler.SearchHandler)
	e.POST("/newtweet", handler.PostTweetHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)

}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/oribe1115/practice-twitter/handler"
	"github.com/oribe1115/practice-twitter/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})

	// dbのテスト用
	e.GET("create/table", handler.CreateTableHandler)

	// e.GET("/authorize", handler.GetRequestTokenHandler)
	e.GET("/authorize/callback", handler.GetAccessTokenHandler)

	e.GET("/search", handler.SearchHandler)
	// e.POST("/newtweet", handler.PostTweetHandler)

	// withTwitter := e.Group("")
	// withTwitter.Use(handler.CheckAuthorization)
	// withTwitter.POST("/newtweet", handler.PostTweetHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)

}

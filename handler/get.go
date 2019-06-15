package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/practice-twitter/model"
)

func SearchHandler(c echo.Context) error {
	tweets, err := model.Search("golang")

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to search")
	}

	return c.JSON(http.StatusOK, tweets)
}

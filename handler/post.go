package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/oribe1115/practice-twitter/model"
)

func PostTweetHandler(c echo.Context) error {
	tweet, err := model.PostTweet()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to post tweet")
	}

	return c.JSON(http.StatusCreated, tweet)
}

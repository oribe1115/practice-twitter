package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"

	"github.com/oribe1115/practice-twitter/model"
)

var (
	credential            *oauth.Credentials
	apiInHandler          *anaconda.TwitterApi
	apiInHandlerWithToken *anaconda.TwitterApi
)

// あとで消す
func GetRequestTokenHandler(c echo.Context) error {
	apiInHandler = anaconda.NewTwitterApiWithCredentials("", "", os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	url, tmpCred, err := apiInHandler.AuthorizationURL(os.Getenv("CALLBACK_URL"))

	fmt.Println(url)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to send authorizeing request")
	}
	credential = tmpCred

	fmt.Println("success to send authorizeing request")
	return c.String(http.StatusOK, "success to send authorizeing request")
}

func GetAccessTokenHandler(c echo.Context) error {
	verifier := c.QueryParam("oauth_verifier")

	tmpCred, _, err := apiInHandler.GetCredentials(credential, verifier)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get access token")
	}
	apiInHandlerWithToken = anaconda.NewTwitterApiWithCredentials(tmpCred.Token, tmpCred.Secret, os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	model.SetAPI(apiInHandlerWithToken)

	fmt.Println("success to get access token")

	tokenCookie := new(http.Cookie)
	tokenCookie.Name = "Token"
	tokenCookie.Value = tmpCred.Token
	c.SetCookie(tokenCookie)

	secretCookie := new(http.Cookie)
	secretCookie.Name = "Secret"
	secretCookie.Value = tmpCred.Secret
	c.SetCookie(secretCookie)

	fmt.Println("success to set Cookie")

	return c.String(http.StatusOK, "success to get access token")
}

func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if apiInHandler == nil {
			protoAPI, tmpCred, url, err := model.GetRequestToken()
			if err != nil {
				fmt.Println(err)
				return c.String(http.StatusInternalServerError, "faild to get request token")
			}
			apiInHandler = protoAPI
			credential = tmpCred
			return c.String(http.StatusOK, url)
		}

		if apiInHandlerWithToken == nil {
			tokenCookie, err := c.Cookie("Token")
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusInternalServerError, "faild to get tokenCookie")
			}
			secretCookie, err := c.Cookie("Secret")
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusInternalServerError, "faild to get secretCookie")
			}

			apiInHandlerWithToken = anaconda.NewTwitterApiWithCredentials(tokenCookie.Value, secretCookie.Value, os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
			model.SetAPI(apiInHandlerWithToken)
		}
		return next(c)
	}
}

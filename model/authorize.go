package model

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

func GetRequestToken() (*anaconda.TwitterApi, *oauth.Credentials, string, error) {
	protoAPI := anaconda.NewTwitterApiWithCredentials("", "", os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	url, tmpCred, err := protoAPI.AuthorizationURL(os.Getenv("CALLBACK_URL"))

	if err != nil {
		return nil, nil, "", err
	}

	return protoAPI, tmpCred, url, nil

}

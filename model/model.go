package model

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

var (
	api          *anaconda.TwitterApi
	credentional *oauth.Credentials
)

// // Twitter Apiのためのkeyなどをセットする
// func GetTwitterAPI() {
// 	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
// 	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
// 	api = anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
// }

// // アクセストークンを取得する
// func GetRequestToken() error {
// 	api = anaconda.NewTwitterApiWithCredentials("", "", os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
// 	_, tmpCred, err := api.AuthorizationURL(os.Getenv("CALLBACK_URL"))
// 	if err != nil {
// 		return err
// 	}
// 	credentional = tmpCred
// 	return nil
// }

func SetAPI(apiInHandler *anaconda.TwitterApi) {
	api = apiInHandler
}

package handler

import (
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/oribe1115/practice-twitter/model"
)

var (
	credential   *oauth.Credentials
	apiInHandler *anaconda.TwitterApi
)

func GetRequestTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	http.DefaultClient.Transport = &urlfetch.Transport{Context: ctx}

	anaconda.SetConsumerKey("consumerKey")
	anaconda.SetConsumerSecret("consumerSecret")

	url, tmpCred, err := apiInHandler.AuthorizationURL(os.Getenv("CALLBACK_URL"))
	if err != nil {
		return
	}
	credential = tmpCred
	http.Redirect(w, r, url, http.StatusFound)
}

func GetAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	http.DefaultClient.Transport = &urlfetch.Transport{Context: ctx}

	c, _, err := apiInHandler.GetCredentials(credential, r.URL.Query().Get("oauth_verifier"))
	if err != nil {
		return
	}
	apiInHandler = anaconda.NewTwitterApi(c.Token, c.Secret)
	model.SetAPI(apiInHandler)
	// TEST POST
	model.PostTweet()
}

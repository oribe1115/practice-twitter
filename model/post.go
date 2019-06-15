package model

import (
	"github.com/ChimeraCoder/anaconda"
)

func PostTweet() (anaconda.Tweet, error) {
	tweet, err := api.PostTweet("test from api", nil)

	if err != nil {
		return tweet, err
	}

	return tweet, nil
}
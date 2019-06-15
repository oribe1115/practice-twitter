package model

import (
	"net/url"
)

func Search(searchWord string) ([]string, error) {
	v := url.Values{}
	v.Set("count", "10")

	searchResult, err := api.GetSearch(searchWord, v)
	if err != nil {
		return nil, err
	}

	tweets := searchResult.Statuses

	var tweetTexts []string

	for _, tweet := range tweets {
		tweetTexts = append(tweetTexts, tweet.Text)
	}

	return tweetTexts, nil

}

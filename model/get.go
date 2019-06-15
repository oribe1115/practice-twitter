package model

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func Search(searchWord string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("count", "10")

	searchResult, err := api.GetSearch(searchWord, v)
	if err != nil {
		return nil, err
	}

	tweets := searchResult.Statuses

	return tweets, nil

}

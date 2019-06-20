package model

import (
	"errors"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jinzhu/gorm"
)

var (
	api *anaconda.TwitterApi
	db  *gorm.DB
)

func SetAPI(apiInHandler *anaconda.TwitterApi) {
	api = apiInHandler
}

func EstablishConnection() (*gorm.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	_db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, errors.New("faild to connect to DB")
	}
	db = _db

	return db, nil
}

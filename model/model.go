package model

import (
	"errors"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	api          *anaconda.TwitterApi
	credentional *oauth.Credentials
	db           *gorm.DB
)

// .envファイルを読み込む
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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

package main

import (
	"github.com/oribe1115/practice-twitter/model"
)

func main() {
	model.LoadEnv()
	model.GetTwitterAPI()

}

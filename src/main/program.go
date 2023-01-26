package main

import (
	"log"

	"github.com/src/main/app"
	_ "github.com/src/resources/docs"
)

// @title Golang Template API
// @description This is a sample golang template api. Have fun.
// @host go-fiber-app.herokuapp.com
// @basePath /
// @version v1.
func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"

	"github.com/MarceloBxD/goopportunities/config"
	"github.com/MarceloBxD/goopportunities/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	// initalize router
	router.Initialize()
}

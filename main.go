package main

import (
	"catchreview-api-app/config"
	"fmt"
	"log"
)

func main() {

	cfg, err := config.ConfInitialize()
	if err != nil {
		log.Fatalln("[main] failed config initialize err : ", err)
		return
	}

	fmt.Println("api port : ", cfg.ApiPort)
}

package main

import (
	"log"

	"github.com/coobeet/onlineboutique/src/currencyservice/service"
)

func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

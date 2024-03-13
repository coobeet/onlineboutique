package main

import (
	"log"

	"github.com/coobeet/onlineboutique/src/cartservice/service"
)

func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

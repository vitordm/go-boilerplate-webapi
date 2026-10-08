package main

import (
	"log"

	"github.com/vitordm/go-boilerplate-webapi/internal/composition"
)

func main() {
	app := composition.NewAPI()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

}

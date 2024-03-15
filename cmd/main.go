package main

import (
	filmlib "github/film-lib"
	"log"
)

func main() {
	srv := new(filmlib.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s\n", err.Error())
	}
}

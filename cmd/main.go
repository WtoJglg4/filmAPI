package main

import (
	filmlib "github/film-lib"
	handler "github/film-lib/pkg/handlers"
	"log"
)

func main() {
	srv := new(filmlib.Server)
	mux := handler.NewHandler()

	if err := srv.Run("8080", mux.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s\n", err.Error())
	}
}

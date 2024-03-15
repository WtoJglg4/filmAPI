package main

import (
	filmlib "github/film-lib"
	handler "github/film-lib/pkg/handler"
	"github/film-lib/pkg/repository"
	"github/film-lib/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	mux := handler.NewHandler(services)
	srv := new(filmlib.Server)

	if err := srv.Run("8080", mux.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s\n", err.Error())
	}
}

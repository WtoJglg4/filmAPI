package handler

import (
	"fmt"
	"github/film-lib/pkg/service"
	"net/http"

	_ "github/film-lib/docs"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", viper.GetString("port"))),
	))

	mux.HandleFunc("/auth/sign-up/", h.signUp)
	mux.HandleFunc("/auth/sign-in/", h.signIn)

	mux.HandleFunc("/actors/", h.actors)
	mux.HandleFunc("/films/", h.films)
	mux.HandleFunc("/films/by-part/", h.filmsByPart)

	return mux
}

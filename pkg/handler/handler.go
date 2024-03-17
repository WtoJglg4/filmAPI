package handler

import (
	"github/film-lib/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-up/", h.signUp)
	mux.HandleFunc("/auth/sign-in/", h.signIn)

	mux.HandleFunc("/actors/", h.actors)
	mux.HandleFunc("/films/", h.films)
	mux.HandleFunc("/films/by-part/", h.filmsByPart)

	return mux
}

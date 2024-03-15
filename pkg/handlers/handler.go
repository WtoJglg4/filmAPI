package handler

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-up/", h.signUp)
	mux.HandleFunc("/auth/sign-in/", h.signIn)

	mux.HandleFunc("/actors/", h.actors)
	mux.HandleFunc("/films/", h.films)
	mux.HandleFunc("/films/by-name/", h.filmsByName)
	mux.HandleFunc("/films/by-actor/", h.filmsByActor)

	return mux
}

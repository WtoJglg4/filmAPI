package handler

import "net/http"

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sign-up"))
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sign-In"))
}

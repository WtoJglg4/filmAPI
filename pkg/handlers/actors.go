package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) actors(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Query().Get("id"))
}

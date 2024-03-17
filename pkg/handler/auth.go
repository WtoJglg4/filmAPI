package handler

import (
	"encoding/json"
	filmapi "github/film-lib"
	"net/http"

	"github.com/sirupsen/logrus"
)

// @Summary		signUp
// @Tags			auth
// @Description	create account
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		filmapi.User	true	"account info"
// @Success		200		{object}	integer			"id"
// @Failure		400,405	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/auth/sign-up/ [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	logrus.Println(r.Method, r.URL.Path)
	w.Header().Set("Content-type", "application/json")

	if r.Method != "POST" {
		newErrorResponse(w, http.StatusMethodNotAllowed, "error: allowed only POST method")
		return
	}

	var input filmapi.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Username == "" || input.Password == "" {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have all the required fields")
		return
	}
	input.Role = "default"
	defer r.Body.Close()

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

// @Summary		signIn
// @Tags			auth
// @Description	login
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input		body		filmapi.User	true	"credentials"
// @Success		200			{string}	string			"token"
// @Failure		400,404,405	{object}	errorResponse
// @Failure		500			{object}	errorResponse
// @Failure		default		{object}	errorResponse
// @Router			/auth/sign-in/ [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	logrus.Println(r.Method, r.URL.Path)
	w.Header().Set("Content-type", "application/json")

	if r.Method != "POST" {
		newErrorResponse(w, http.StatusMethodNotAllowed, "error: allowed only POST method")
		return
	}

	var input filmapi.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Username == "" || input.Password == "" {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have all the required fields")
		return
	}
	input.Role = "default"
	defer r.Body.Close()

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}

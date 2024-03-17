package handler

import (
	"encoding/json"
	filmapi "github/film-lib"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (h *Handler) actors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	logrus.Println(r.Method, r.URL)
	err := h.userIdentity(w, r)
	if err != nil {
		return
	}

	param := r.URL.Query().Get("id")
	switch param {
	case "":
		switch r.Method {
		case "GET":
			h.GetActorsList(w, r)
		case "POST":
			h.CreateActor(w, r)
		default:
			newErrorResponse(w, http.StatusMethodNotAllowed, "error: method not allowed")
			return
		}
	default:
		actorId, err := strconv.Atoi(param)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "error: invalid id parameter")
			return
		}

		switch r.Method {
		case "GET":
			h.GetActorById(w, r, actorId)
		case "PUT":
			h.UpdateActorById(w, r, actorId)
		case "DELETE":
			h.DeleteActorById(w, r, actorId)
		default:
			newErrorResponse(w, http.StatusMethodNotAllowed, "error: method not allowed")
			return
		}
	}
}

func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	var input filmapi.Actor

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Name == "" || input.Gender == "" || input.BirthDate == "" {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have all the required fields")
		return
	}
	defer r.Body.Close()

	id, err := h.services.CreateActor(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GetActorsList(w http.ResponseWriter, r *http.Request) {
	_, err := getUserId(w)
	if err != nil {
		return
	}

	actors_list, err := h.services.GetActorsList()
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(actors_list)
}

func (h *Handler) GetActorById(w http.ResponseWriter, r *http.Request, id int) {
	_, err := getUserId(w)
	if err != nil {
		return
	}

	actor, err := h.services.GetActorById(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(actor)
}

func (h *Handler) UpdateActorById(w http.ResponseWriter, r *http.Request, id int) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	var input filmapi.Actor

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Name == "" && input.Gender == "" && input.BirthDate == "" {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have fields")
		return
	}
	defer r.Body.Close()

	err = h.services.UpdateActorById(input.Name, input.Gender, input.BirthDate, id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteActorById(w http.ResponseWriter, r *http.Request, id int) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	err = h.services.DeleteActorById(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}

package handler

import (
	"encoding/json"
	filmapi "github/film-lib"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (h *Handler) films(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	logrus.Println(r.Method, r.URL)
	err := h.userIdentity(w, r)
	if err != nil {
		return
	}
	sort_params := map[string]bool{
		"name":         true,
		"rating":       true,
		"release_date": true,
	}

	param_id := r.URL.Query().Get("id")
	param_sort := r.URL.Query().Get("sort")

	switch r.Method {
	case "GET":
		if param_sort == "" {
			param_sort = "rating"
		} else {
			if _, ok := sort_params[param_sort]; !ok {
				newErrorResponse(w, http.StatusBadRequest, "error: invalid sort parameter")
				return
			}
		}
		h.GetFilmsList(w, r, param_sort)
	case "POST":
		h.CreateFilm(w, r)
	case "PUT":
		filmId, err := strconv.Atoi(param_id)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "error: invalid id parameter")
			return
		}
		h.UpdateFilmById(w, r, filmId)
	case "DELETE":
		filmId, err := strconv.Atoi(param_id)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "error: invalid id parameter")
			return
		}
		h.DeleteFilmById(w, r, filmId)
	default:
		newErrorResponse(w, http.StatusMethodNotAllowed, "error: method not allowed")
		return
	}
}

func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	var input filmapi.Film

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Name == "" || input.Description == "" || input.ReleaseDate == "" || input.Rating <= 0 {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have all the required fields")
		return
	}
	defer r.Body.Close()

	id, err := h.services.CreateFilm(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetFilmsList(w http.ResponseWriter, r *http.Request, sort string) {
	_, err := getUserId(w)
	if err != nil {
		return
	}

	films_list, err := h.services.GetFilmsList(sort)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(films_list)
}

func (h *Handler) UpdateFilmById(w http.ResponseWriter, r *http.Request, id int) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	var input filmapi.Film

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	} else if input.Name == "" && input.Description == "" && input.ReleaseDate == "" && input.Rating == 0 {
		newErrorResponse(w, http.StatusBadRequest, "error: passed structure does not have fields")
		return
	}
	defer r.Body.Close()

	err = h.services.UpdateFilmById(input.Name, input.Description, input.ReleaseDate, input.Rating, id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteFilmById(w http.ResponseWriter, r *http.Request, id int) {
	_, err := getUserId(w)
	if err != nil {
		return
	}
	_, err = getUserRole(w)
	if err != nil {
		return
	}

	err = h.services.DeleteFilmById(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) filmsByPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	logrus.Println(r.Method, r.URL)
	err := h.userIdentity(w, r)
	if err != nil {
		return
	}

	param_actor := r.URL.Query().Get("actor")
	param_name := r.URL.Query().Get("name")

	switch r.Method {
	case "GET":
		if param_actor != "" {
			h.GetFilmByPart(w, r, "actor", param_actor)
		} else if param_name != "" {
			h.GetFilmByPart(w, r, "name", param_name)
		} else {
			newErrorResponse(w, http.StatusBadRequest, "error: not enough parameters")
		}
	default:
		newErrorResponse(w, http.StatusMethodNotAllowed, "error: method not allowed")
	}
}

func (h *Handler) GetFilmByPart(w http.ResponseWriter, r *http.Request, parameter, req string) {
	_, err := getUserId(w)
	if err != nil {
		return
	}

	films_list, err := h.services.GetFilmByPart(parameter, req)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(films_list)
}

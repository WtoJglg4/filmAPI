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

// @Summary		Create film
// @Security		ApiKeyAuth
// @Tags			films
// @Description	create film
// @ID				create-film
// @Accept			json
// @Produce		json
// @Param			input	body		filmapi.Film	true	"film info"
// @Success		200		{object}	integer			"id"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/films/ [post]
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

// @Summary		Get films list
// @Security		ApiKeyAuth
// @Tags			films
// @Description	get films list sorted by name, rating or release_date (default by rating descending)
// @ID				get-films-list
// @Produce		json
// @Param			sort	query		string	false	"Sorting parameter: "name", "rating", "release_date" (default by rating descending)"
// @Success		200		{array}		filmapi.Film
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/films/ [get]
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

// @Summary		Update film by id
// @Security		ApiKeyAuth
// @Tags			films
// @Description	update film by id
// @ID				update-film-by-id
// @Accept			json
// @Produce		json
// @Param			input	body		filmapi.Film	true	"film info"
// @Param			id		query		int				true	"Film ID"
// @Success		200		{string}	string			"status"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/films/ [put]
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

// @Summary		Delete film by id
// @Security		ApiKeyAuth
// @Tags			films
// @Description	delete film by id
// @ID				delete-film-by-id
// @Produce		json
// @Param			id		query		int		true	"Film ID"
// @Success		200		{string}	string	"status"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/films/ [delete]
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

// @Summary		Get film by part
// @Security		ApiKeyAuth
// @Tags			films
// @Description	Get film by a fragment of its title or actor's name
// @ID				get-film-by-part
// @Produce		json
// @Param			name	query		string	false	"film name"
// @Param			actor	query		string	false	"actor`s name"
// @Success		200		{array}		filmapi.Film
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/films/by-part/{param} [get]
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

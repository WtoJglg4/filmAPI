package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	userRoleCtx         = "userRole"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) error {
	header := r.Header[authorizationHeader]
	if len(header) == 0 {
		newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
		return errors.New("empty auth header")
	} else if len(header) > 1 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return errors.New("invalid auth header")
	}

	headerParts := strings.Split(header[0], " ")
	if len(headerParts) != 2 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return errors.New("invalid auth header")
	}

	userId, userRole, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return errors.New("invalid auth header")
	}

	w.Header().Set(userIdCtx, fmt.Sprint(userId))
	w.Header().Set(userRoleCtx, fmt.Sprint(userRole))
	return nil
}

func getUserId(w http.ResponseWriter) (int, error) {
	userId := w.Header().Get(userIdCtx)
	id, err := strconv.Atoi(userId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "user id is not found")
		logrus.Error("user id is not found")
		return 0, errors.New("user id is not found")
	}
	return id, nil
}

func getUserRole(w http.ResponseWriter) (string, error) {
	userRole := w.Header().Get(userRoleCtx)
	if userRole != "admin" {
		newErrorResponse(w, http.StatusForbidden, "not enough rights")
		logrus.Error("not enough rights")
		return "", errors.New("not enough rights")
	}
	return userRole, nil
}

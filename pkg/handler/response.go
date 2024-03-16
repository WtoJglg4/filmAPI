package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	bytes, _ := json.Marshal(Error{Message: message})
	http.Error(w, fmt.Sprint(string(bytes)), statusCode)
}

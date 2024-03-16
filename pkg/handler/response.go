package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type error struct {
	Message string
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	bytes, _ := json.Marshal(error{Message: message})
	http.Error(w, fmt.Sprint(string(bytes)), statusCode)
}

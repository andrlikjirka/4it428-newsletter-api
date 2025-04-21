package utils

import (
	"4it428-newsletter-api/libs/logger"
	"encoding/json"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, body any) {
	logger.Init()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(body)
	_, err := w.Write(b)
	if err != nil {
		logger.Error("Error writing response.", "error", err)
	}
}

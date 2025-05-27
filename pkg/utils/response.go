package utils

import (
	"4it428-newsletter-api/libs/logger"
	"encoding/json"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, statusCode int, err error) {
	logger.Init()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err == nil {
		return
	}

	resp := map[string]string{
		"error": err.Error(),
	}

	b, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		logger.Error("Failed to marshal error response", "error", marshalErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(b)
	if writeErr != nil {
		logger.Error("Failed to write error response", "error", writeErr)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, body any) {
	logger.Init()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || body == nil {
		return // Don't write a body for 204 or if body is nil
	}

	b, err := json.Marshal(body)
	if err != nil {
		logger.Error("Error marshaling response body.", "error", err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		logger.Error("Error writing response.", "error", err)
	}
}

package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Success sends a JSON success response with status code and headers
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	marshaled, err := json.Marshal(data)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(marshaled)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
	}
}

// Error sends a JSON error response with status code and headers
func Error(w http.ResponseWriter, statusCode int, err error) {
	resp := errorResponse{Success: false, Error: err.Error()}
	marshaled, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(marshaled)
	if err != nil {
		log.Println(err)
	}
}

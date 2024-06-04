package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func ApiErr(w http.ResponseWriter, status string, message string, data map[string]string, code int) {
	resp := ErrResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func ApiInternalErr(w http.ResponseWriter, err error) {
	log.Println("apiInternalErr:", err)
	ApiErr(w, "error", "server error", nil, http.StatusInternalServerError)
}

func ApiBadRequestErr(w http.ResponseWriter, message string, data map[string]string) {
	ApiErr(w, "fail", message, data, http.StatusBadRequest)
}

func ApiUnauthorizedErr(w http.ResponseWriter, message string, data map[string]string) {
	ApiErr(w, "fail", message, data, http.StatusUnauthorized)
}

func ApiNotFoundErr(w http.ResponseWriter, message string) {
	ApiErr(w, "fail", message, nil, http.StatusNotFound)
}

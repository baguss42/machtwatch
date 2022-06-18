package handler

import (
	"encoding/json"
	"net/http"
)

const ErrMethodNotAllowed = "Method Not Allowed"

type ResponsePayload struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta map[string]int

func NewResponse(data interface{}, httpCode int) *ResponsePayload {
	return &ResponsePayload{
		Data: data,
		Meta: map[string]int{
			"meta": httpCode,
		},
	}
}

func WriteSuccess(w http.ResponseWriter, payload interface{}) (int, error) {
	resp := NewResponse(payload, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return http.StatusOK, nil
}

func WriteCreated(w http.ResponseWriter, payload interface{}) (int, error) {
	resp := NewResponse(payload, http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	return http.StatusCreated, nil
}

func WriteErrorMethodNotAllowed(w http.ResponseWriter, err error) (int, error) {
	resp := NewResponse(err.Error(), http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(resp)
	return http.StatusMethodNotAllowed, err
}

func WriteInternalServerError(w http.ResponseWriter, err error) (int, error) {
	resp := NewResponse(err.Error(), http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
	return http.StatusInternalServerError, err
}

func WriteErrorBadRequest(w http.ResponseWriter, err error) (int, error) {
	resp := NewResponse(err.Error(), http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
	return http.StatusBadRequest, err
}

func WriteErrorNotFound(w http.ResponseWriter, err error) (int, error) {
	resp := NewResponse(err.Error(), http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(resp)
	return http.StatusNotFound, err
}
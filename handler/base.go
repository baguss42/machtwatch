package handler

import (
	"encoding/json"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"net/http"
)

const ErrMethodNotAllowed = "Method Not Allowed"

type ResponsePayload struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta map[string]int

func NewPayload(data interface{}, httpCode int) *ResponsePayload {
	return &ResponsePayload{
		Data: data,
		Meta: map[string]int{
			"meta": httpCode,
		},
	}
}

type CustomResponse struct {
	Result      interface{}
	CustomError entity.CustomError
}

func NewResponse() CustomResponse {
	return CustomResponse{
		Result:      nil,
		CustomError: entity.NewCustomError(),
	}
}

func (c *CustomResponse) Write(w http.ResponseWriter) (int, error) {
	if c.Result == nil {
		c.Result = "ok" // ok by default
	}
	response := NewPayload(c.Result, c.CustomError.HttpCode)
	if c.CustomError.Err != nil {
		response = NewPayload(c.CustomError.Err.Error(), c.CustomError.HttpCode)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c.CustomError.HttpCode)
	_ = json.NewEncoder(w).Encode(response)
	return c.CustomError.HttpCode, c.CustomError.Err
}

func (c *CustomResponse) ErrorMethodNotAllowed(w http.ResponseWriter) (int, error) {
	c.CustomError.HttpCode = http.StatusMethodNotAllowed
	c.CustomError.Err = errors.New(ErrMethodNotAllowed)
	return c.Write(w)
}

func (c *CustomResponse) ErrorBadRequest(w http.ResponseWriter, err error) (int, error) {
	c.CustomError.HttpCode = http.StatusBadRequest
	if err != nil {
		c.CustomError.Err = err
	}
	return c.Write(w)
}

func (c *CustomResponse) ErrorRecordNotFound(w http.ResponseWriter, err error) (int, error) {
	c.CustomError.HttpCode = http.StatusNotFound
	if err != nil {
		c.CustomError.Err = err
	}
	return c.Write(w)
}

func (c *CustomResponse) ErrorInternalServerError(w http.ResponseWriter, err error) (int, error) {
	c.CustomError.HttpCode = http.StatusInternalServerError
	if err != nil {
		c.CustomError.Err = err
	}
	return c.Write(w)
}

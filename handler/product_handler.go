package handler

import (
	"encoding/json"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/service"
	"io"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service service.ProductServiceInterface
}

func DecodeProduct(r io.Reader) (entity.Product, error) {
	var Product entity.Product
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&Product); err != nil {
		return Product, errors.New("could not decode body")
	}

	return Product, nil
}

func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()
	switch r.Method {
	case http.MethodPost:
		return h.Create(w, r)
	case http.MethodGet:
		return h.Get(w, r)
	default:
		return response.ErrorMethodNotAllowed(w)
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()
	var product entity.Product

	product, response.CustomError.Err = DecodeProduct(r.Body)
	if response.CustomError.Err != nil {
		return response.ErrorBadRequest(w, errors.New("request body are invalid"))
	}

	response.CustomError = h.Service.Create(r.Context(), product)
	if response.CustomError.Err == nil {
		response.CustomError.HttpCode = http.StatusCreated
	}

	return response.Write(w)
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()
	var id int64

	idParam := r.URL.Query().Get("id")
	id, response.CustomError.Err = strconv.ParseInt(idParam, 10, 64)
	if response.CustomError.Err != nil || id < 1 {
		return response.ErrorBadRequest(w, errors.New("id is not valid"))
	}

	response.Result, response.CustomError = h.Service.Get(r.Context(), id)
	return response.Write(w)
}

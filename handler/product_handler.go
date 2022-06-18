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
	switch r.Method {
	case http.MethodPost:
		return h.Create(w,r)
	case http.MethodGet:
		return h.Get(w,r)
	default:
		return WriteErrorMethodNotAllowed(w, errors.New(ErrMethodNotAllowed))
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	Product, err := DecodeProduct(r.Body)
	if err != nil {
		return WriteErrorBadRequest(w, err)
	}
	err = h.Service.Create(r.Context(), Product)
	if err != nil {// TODO check if error is bad request/internal server error
		return WriteInternalServerError(w, err)
	}
	return WriteCreated(w, Product)
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) (int, error) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return WriteErrorBadRequest(w, err)
	}

	result, err := h.Service.Get(r.Context(), id)
	if err != nil {// TODO check if error is bad request/internal server error
		if err.Error() == "sql: no rows in result set" {
			return WriteErrorNotFound(w, errors.New("record is not found"))
		}
		return WriteInternalServerError(w, err)
	}

	return WriteSuccess(w, result)
}
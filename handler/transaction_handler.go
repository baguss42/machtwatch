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

type TransactionHandler struct {
	Service service.TransactionServiceInterface
}

func DecodeTransactionOrder(r io.Reader) (entity.TransactionOrder, error) {
	var transactionOrder entity.TransactionOrder
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&transactionOrder); err != nil {
		return transactionOrder, errors.New("could not decode body")
	}

	return transactionOrder, nil
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()
	switch r.Method {
	case http.MethodPost:
		return h.Create(w,r)
	case http.MethodGet:
		return h.Get(w,r)
	default:
		return response.ErrorMethodNotAllowed(w)
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()

	var transactionOrder entity.TransactionOrder
	transactionOrder, response.CustomError.Err = DecodeTransactionOrder(r.Body)
	if response.CustomError.Err != nil || len(transactionOrder.Carts) < 1 {
		return response.ErrorBadRequest(w, errors.New("invalid body request"))
	}

	response.CustomError = h.Service.Create(r.Context(), transactionOrder)

	return response.Write(w)
}

func (h *TransactionHandler) Get(w http.ResponseWriter, r *http.Request) (int, error) {
	response := NewResponse()
	var id int64

	idParam := r.URL.Query().Get("id")
	id, response.CustomError.Err = strconv.ParseInt(idParam, 10, 64)
	if response.CustomError.Err != nil {
		return response.ErrorBadRequest(w, errors.New("id is invalid"))
	}

	response.Result, response.CustomError = h.Service.Get(r.Context(), id)

	return response.Write(w)
}
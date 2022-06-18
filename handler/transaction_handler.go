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
	switch r.Method {
	case http.MethodPost:
		return h.Create(w,r)
	case http.MethodGet:
		return h.Get(w,r)
	default:
		return WriteErrorMethodNotAllowed(w, errors.New(ErrMethodNotAllowed))
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	transactionOrder, err := DecodeTransactionOrder(r.Body)
	if err != nil || len(transactionOrder.Carts) < 1 {
		err = errors.New("invalid body request")
		return WriteErrorBadRequest(w, err)
	}

	err = h.Service.Create(r.Context(), transactionOrder)
	if err != nil {// TODO check if error is bad request/internal server error
		return WriteInternalServerError(w, err)
	}

	return WriteCreated(w, "ok")
}

func (h *TransactionHandler) Get(w http.ResponseWriter, r *http.Request) (int, error) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return WriteErrorBadRequest(w, err)
	}

	result, err := h.Service.Get(r.Context(), id)
	if err != nil {// TODO check if error is bad request/internal server error
		return WriteInternalServerError(w, err)
	}
	if len(result) < 1 {
		return WriteErrorNotFound(w, errors.New("records are not found"))
	}

	return WriteSuccess(w, result)
}
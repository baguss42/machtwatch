package handler

import (
	"encoding/json"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/service"
	"io"
	"net/http"
)

type BrandHandler struct {
	Service service.BrandServiceInterface
}

func DecodeBrand(r io.Reader) (entity.Brand, error) {
	var brand entity.Brand
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&brand); err != nil {
		return brand, errors.New("could not decode body")
	}

	return brand, nil
}

func (h *BrandHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		return WriteErrorMethodNotAllowed(w, errors.New(ErrMethodNotAllowed))
	}

	brand, err := DecodeBrand(r.Body)
	if err != nil {
		return WriteErrorBadRequest(w, err)
	}

	err = h.Service.Create(r.Context(), brand)
	if err != nil {// TODO check if error is bad request/internal server error
		return WriteInternalServerError(w, err)
	}

	return WriteCreated(w, brand)
}
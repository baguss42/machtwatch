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
	response := NewResponse()
	if r.Method != http.MethodPost {
		return response.ErrorMethodNotAllowed(w)
	}

	var brand entity.Brand
	brand, response.CustomError.Err = DecodeBrand(r.Body)
	if response.CustomError.Err != nil {
		return response.ErrorBadRequest(w, nil)
	}

	response.CustomError = h.Service.Create(r.Context(), brand)
	if response.CustomError.Err == nil {
		response.CustomError.HttpCode = http.StatusCreated
	}
	return response.Write(w)
}

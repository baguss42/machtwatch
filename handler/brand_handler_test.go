package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/test/factory"
	mock "github.com/baguss42/machtwatch/test/mock/service"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BrandHandlerTestSuite struct {
	suite.Suite
	mockService     *mock.BrandServiceMock
	handlerInstance BrandHandler
	brand           entity.Brand
	customError     entity.CustomError
}

func (ts *BrandHandlerTestSuite) SetupTest() {
	ts.mockService = &mock.BrandServiceMock{}
	ts.handlerInstance = BrandHandler{Service: ts.mockService}
	ts.brand = (&factory.BrandFactory{}).Build()
	ts.customError = entity.NewCustomError()
}

func (ts *BrandHandlerTestSuite) TestDecodeBrand() {
	brandByte, _ := json.Marshal(ts.brand)
	bodyRequest := bytes.NewReader(brandByte)
	brand, err := DecodeBrand(bodyRequest)
	ts.Equal(brand, ts.brand)
	ts.Nil(err)

	bodyRequest = bytes.NewReader([]byte("abc"))
	brand, err = DecodeBrand(bodyRequest)
	ts.NotEqual(brand, ts.brand)
	ts.NotNil(err)
}

func (ts *BrandHandlerTestSuite) TestCreate() {
	testCases := []struct {
		name          string
		brand         entity.Brand
		method        string
		err           entity.CustomError
		isErrorDecode bool
		httpStatus    int
	}{
		{
			name:       "Test case 1: success create brand | 200",
			brand:      ts.brand,
			method:     http.MethodPost,
			err:        ts.customError,
			httpStatus: http.StatusCreated,
		},
		{
			name:       "Test case 2: failed method not allowed | 405",
			brand:      ts.brand,
			method:     http.MethodGet,
			err:        ts.customError,
			httpStatus: http.StatusMethodNotAllowed,
		},
		{
			name:          "Test case 3: failed decode brand | 400",
			brand:         entity.Brand{},
			method:        http.MethodPost,
			isErrorDecode: true,
			err:           entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("could not decode body")},
			httpStatus:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		brandByte, _ := json.Marshal(tc.brand)
		if tc.isErrorDecode {
			brandByte, _ = json.Marshal([]byte(""))
		}
		r := httptest.NewRequest(tc.method, "/brand", bytes.NewReader(brandByte))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ts.mockService.On("Create", r.Context(), tc.brand).Return(tc.err).Once()

		_, err := ts.handlerInstance.Create(w, r)
		res := w.Result()
		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			ts.Error(err, "unexpected error")
		}

		ts.Equal(tc.httpStatus, res.StatusCode)

		res.Body.Close()
	}
}

func TestBrandHandler(t *testing.T) {
	suite.Run(t, new(BrandHandlerTestSuite))
}

package handler

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestErrorRecordNotFound(t *testing.T) {
	c := NewResponse()
	w := httptest.NewRecorder()
	code, err := c.ErrorRecordNotFound(w, errors.New("record not exist"))
	if code != 404 && err == nil {
		t.Error(fmt.Sprintf("expected code: 404, actual code: %d | expected error is not nil, actual is nil", code))
	}
}

func TestErrorInternalServerError(t *testing.T) {
	c := NewResponse()
	w := httptest.NewRecorder()
	code, err := c.ErrorInternalServerError(w, errors.New("unexpected error"))
	if code != 500 && err == nil {
		t.Error(fmt.Sprintf("expected code: 500, actual code: %d | expected error is not nil, actual is nil", code))
	}
}

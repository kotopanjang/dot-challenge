package wrapper_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"dot-rahadian-ardya-kotopanjang/pkg/wrapper"

	"github.com/stretchr/testify/assert"
)

func TestResponseSuccess(t *testing.T) {
	t.Run("should return success json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Data = "OK"
		result.Status = wrapper.StatOK

		wrapper.ResponseSuccess(recorder, http.StatusOK, result, "success")

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should return success json response without status declared", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Data = "OK"

		wrapper.ResponseSuccess(recorder, http.StatusOK, result, "success")

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should return error json response with status code 500 caused by invalid assigned status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Data = "OK"

		wrapper.ResponseSuccess(recorder, http.StatusMovedPermanently, result, "success")

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}

func TestResponseError(t *testing.T) {
	t.Run("should return error not found json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Err = wrapper.ErrNotFound
		result.ErrMessage = "Not Found"
		result.Status = wrapper.StatInvalidRequest

		wrapper.ResponseError(recorder, result)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return error bad request json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Err = wrapper.ErrBadRequest
		result.ErrMessage = "Bad Request"
		result.Status = wrapper.StatInvalidRequest

		wrapper.ResponseError(recorder, result)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return error method not allowed json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Err = wrapper.ErrMethodNotAllowed
		result.ErrMessage = "Method Not Allowed"
		result.Status = wrapper.StatInvalidRequest

		wrapper.ResponseError(recorder, result)

		assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
	})

	t.Run("should return error internal server json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Err = errors.New("error")
		result.Status = wrapper.StatUnexpectedError

		wrapper.ResponseError(recorder, result)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}

func TestResultError(t *testing.T) {
	t.Run("should return success result error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		recorder.Code = 404
		result := new(wrapper.Result)
		result.Data = ""
		result.Status = wrapper.StatInvalidRequest

		wrapper.ErrorResult(wrapper.ErrNotFound, "not found", wrapper.StatInvalidRequest)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestResultSuccess(t *testing.T) {
	t.Run("should return success result success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		result := new(wrapper.Result)
		result.Data = ""
		result.Status = wrapper.StatOK

		wrapper.SuccessResult(wrapper.StatOK, "success")

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

package wrapper

import (
	"encoding/json"
	"net/http"
)

// Result is a struct that contains property of result of function's return
type Result struct {
	Err         error
	ErrMessage  string
	Data        interface{}
	PartialMeta *PartialMeta
	Status
	FieldErrors []string
}

// PartialMeta is a struct that contains property to describe the partial result
type PartialMeta struct {
	Page            int `json:"page"`
	TotalPage       int `json:"totalPage"`
	TotalData       int `json:"totalData"`
	TotalDataOnPage int `json:"totalDataOnPage"`
}

type httpResponse struct {
	Success     bool         `json:"success"`
	Data        interface{}  `json:"data"`
	PartialMeta *PartialMeta `json:"meta,omitempty"`
	Message     string       `json:"message"`
	Status      Status       `json:"status"`
	Code        int          `json:"code"`
}

// ResponseSuccess will respond http success with json serializer. But if the status code doesn't represents success, it will be redirected to ResponseError()
//
// The code must be filled with status code 2xx
func ResponseSuccess(w http.ResponseWriter, code int, result *Result, message string) {
	if code < 200 || code >= 300 {
		result.Err = ErrInternalServer
		ResponseError(w, result)
		return
	}
	if result.Status == "" {
		result.Status = StatOK
	}
	responseData := httpResponse{
		Success:     true,
		Data:        result.Data,
		PartialMeta: result.PartialMeta,
		Message:     message,
		Status:      result.Status,
		Code:        code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseData)

}

// ResponseError will respond http error with json serializer
func ResponseError(w http.ResponseWriter, result *Result) {
	code := getHTTPStatusCodeByError(result.Err)
	if result.ErrMessage == "" {
		result.ErrMessage = result.Err.Error()
	}
	responseData := httpResponse{
		Success: false,
		Message: result.ErrMessage,
		Status:  result.Status,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseData)

}

// ErrorResult return result that contain error
func ErrorResult(err error, errMessage string, status Status) Result {
	return Result{
		Err:        err,
		ErrMessage: errMessage,
		Status:     status,
	}
}

// SuccessResult returns success result
func SuccessResult(data interface{}, status Status) Result {
	return Result{
		Data:   data,
		Status: status,
	}
}

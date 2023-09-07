package wrapper

// Status is a type that describe the request
type Status string

// StatOK will describe that the status is currently ok
var StatOK Status = "OK"

// StatInvalidRequest describe that the request is rejected due to invalid parameters, body, or query string
var StatInvalidRequest Status = "INVALID_REQUEST"

// StatUnexpectedError describe that the request is rejected due to invalid parameters, body, or query string
var StatUnexpectedError Status = "UNEXPECTED_ERROR"

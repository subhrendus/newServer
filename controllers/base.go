package controllers

import (
	"net/http"
)

// TODO: build an error glossary

// baseController - the struct to hold base controller
type baseController struct {
	controller
}

// success - base method for sending HTTP 200
func (c *baseController) success(w http.ResponseWriter, r *http.Request, data []byte) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	return err
}

// successEmpty - base method for sending HTTP 200 (without content)
func (c *baseController) successEmpty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// serviceUnavailable - base method for sending HTTP 503
func (c *baseController) serviceUnavailable(w http.ResponseWriter) {
	w.WriteHeader(http.StatusServiceUnavailable)
}

// failure - base method for sending HTTP 500 with additional params
func (c *baseController) failure(w http.ResponseWriter, code int, message string) error {
	// tODO: map against custom error-codes

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(message))
	return err
}

// internalServerError - base method for sending HTTP 500
func (c *baseController) internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

// badRequest - base method for sending HTTP 502
func (c *baseController) badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

// badgateway - base method for sending HTTP 502
func (c *baseController) badGateway(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadGateway)
}

// unauthorized - base method for sending HTTP 401
func (c *baseController) unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

// notimplemented - base method for sending HTTP 503
func (c *baseController) notimplemented(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// forbidden - base method for sending HTTP 403
func (c *baseController) forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

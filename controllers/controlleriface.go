package controllers

import "net/http"

type controller interface {
	badGateway(http.ResponseWriter)
	badRequest(http.ResponseWriter)
	failure(http.ResponseWriter, int, string) error
	forbidden(http.ResponseWriter)
	internalServerError(http.ResponseWriter)
	serviceUnavailable(http.ResponseWriter)
	success(http.ResponseWriter, *http.Request, []byte) error
	successEmpty(http.ResponseWriter)
	unauthorized(http.ResponseWriter)
}

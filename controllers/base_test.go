package controllers

import (
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// test for system controller
var _ = Describe("baseController", func() {
	ctrl := &baseController{}
	var w *httptest.ResponseRecorder

	Context("when any method for the baseController is called", func() {
		BeforeEach(func() {
			w = httptest.NewRecorder()
		})

		It("returns a HTTP 200 when 'SuccessEmpty' is called.", func() {
			ctrl.successEmpty(w)
			Expect(w.Code).To(Equal(200))
		})

		It("returns a HTTP 500 when 'Failure' is called.", func() {
			ctrl.failure(w, 123, "some specific error")
			Expect(w.Code).To(Equal(500))
		})

		It("returns a HTTP 500 when 'InternalServerError' is called.", func() {
			ctrl.internalServerError(w)
			Expect(w.Code).To(Equal(500))
		})

		/*
			XIt("returns a HTTP 502 when 'BadGateway' is called.", func() {
				ctrl.badGateway(w)
				Expect(w.Code).To(Equal(502))
			})
		*/

		It("returns a HTTP 401 when 'Unauthorized' is called.", func() {
			ctrl.unauthorized(w)
			Expect(w.Code).To(Equal(401))
		})

		It("returns a HTTP 503 when 'ServiceUnavailable' is called.", func() {
			ctrl.serviceUnavailable(w)
			Expect(w.Code).To(Equal(503))
		})

		It("returns a HTTP 403 when 'Forbidden' is called.", func() {
			ctrl.forbidden(w)
			Expect(w.Code).To(Equal(403))
		})
	})
})

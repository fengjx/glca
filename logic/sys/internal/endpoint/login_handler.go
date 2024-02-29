package endpoint

import (
	"github.com/fengjx/luchen"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/pb"
	"github.com/fengjx/glca/transport/http"
)

type loginHandler struct {
}

func newLoginHandler() *loginHandler {
	return &loginHandler{}
}

func (h *loginHandler) Bind(router luchen.HTTPRouter) {
	router.Route("/openapi/user", func(r chi.Router) {
		r.Handle("/login", h.login())
	})
}

func (h *loginHandler) login() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeLoginEndpoint(),
		luchen.DecodeJSONRequest[pb.LoginReq],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

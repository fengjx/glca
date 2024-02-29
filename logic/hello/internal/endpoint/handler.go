package endpoint

import (
	"github.com/fengjx/luchen"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/pb"
	"github.com/fengjx/glca/transport/http"
)

type greeterHandler struct {
}

func newGreeterHandler() *greeterHandler {
	return &greeterHandler{}
}

func (h *greeterHandler) Bind(router luchen.HTTPRouter) {
	router.Route("/hello", func(r chi.Router) {
		r.Handle("/say-hello", h.sayHello())
	})
}

func (h *greeterHandler) sayHello() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeSayHelloEndpoint(),
		luchen.DecodeParamHTTPRequest[pb.HelloReq],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

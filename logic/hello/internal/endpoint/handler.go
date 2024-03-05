package endpoint

import (
	"github.com/fengjx/luchen"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/pb"
	"github.com/fengjx/glca/transport/http"
)

type greeterHandler struct {
}

func newGreeterHandler() *greeterHandler {
	return &greeterHandler{}
}

func (h *greeterHandler) Bind(router *luchen.ServeMux) {
	helloAPI := "/hello"
	router.Handle(helloAPI+"/say-hello", h.sayHello())
}

func (h *greeterHandler) sayHello() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeSayHelloEndpoint(),
		luchen.DecodeHTTPParamRequest[pb.HelloReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
		options...,
	)
}

package endpoint

import (
	"github.com/fengjx/luchen"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/pb"
	"github.com/fengjx/glca/transport/http"
)

type loginHandler struct {
}

func newLoginHandler() *loginHandler {
	return &loginHandler{}
}

func (h *loginHandler) Bind(router *luchen.ServeMux) {
	userAPI := http.OpenAPI + "/user"
	router.Handle(userAPI+"/login", h.login())
}

func (h *loginHandler) login() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeLoginEndpoint(),
		luchen.DecodeHTTPJSONRequest[pb.LoginReq],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

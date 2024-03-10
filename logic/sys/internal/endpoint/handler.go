package endpoint

import (
	"github.com/fengjx/luchen"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/protocol"
	"github.com/fengjx/glca/transport/http"
)

type configHandler struct {
}

func newConfigHandler() *configHandler {
	return &configHandler{}
}

func (h configHandler) Bind(router *luchen.ServeMux) {
	router.Handle(http.OpenAPI+"/config/fetch", h.fetch())
}

func (h configHandler) fetch() *httptransport.Server {
	return http.NewHandler(
		MakeConfigFetchEndpoint(),
		luchen.DecodeHTTPJSONRequest[protocol.LoginReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

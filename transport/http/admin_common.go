package http

import (
	"path"

	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/service/common"
)

type adminCommonHandler struct {
}

func newAdminCommonHandler() *adminCommonHandler {
	return &adminCommonHandler{}
}

func (h *adminCommonHandler) Bind(router luchen.HTTPRouter) {
	router.Route(path.Join(adminAPI, "common"), func(r chi.Router) {
		r.Handle("/insert", h.insert())
		r.Handle("/query", h.query())
		r.Handle("/get", h.get())
		r.Handle("/update", h.update())
		r.Handle("/delete", h.delete())
	})
}

func (h *adminCommonHandler) query() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return luchen.NewHTTPHandler(
		common.GetInst().Endpoints.MakeQueryEndpoint(),
		luchen.DecodeJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSON(httpResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) get() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return luchen.NewHTTPHandler(
		common.GetInst().Endpoints.MakeGetEndpoint(),
		luchen.DecodeJSONRequest[daox.GetRecord],
		luchen.EncodeHTTPJSON(httpResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) insert() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return luchen.NewHTTPHandler(
		common.GetInst().Endpoints.MakeInsertEndpoint(),
		luchen.DecodeJSONRequest[daox.InsertRecord],
		luchen.EncodeHTTPJSON(httpResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) update() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return luchen.NewHTTPHandler(
		common.GetInst().Endpoints.MakeUpdateEndpoint(),
		luchen.DecodeJSONRequest[daox.UpdateRecord],
		luchen.EncodeHTTPJSON(httpResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) delete() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return luchen.NewHTTPHandler(
		common.GetInst().Endpoints.MakeDeleteEndpoint(),
		luchen.DecodeJSONRequest[daox.DeleteRecord],
		luchen.EncodeHTTPJSON(httpResponseWrapper),
		options...,
	)
}

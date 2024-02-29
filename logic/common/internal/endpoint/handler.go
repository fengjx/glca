package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/transport/http"
)

type adminCommonHandler struct {
}

func newAdminCommonHandler() *adminCommonHandler {
	return &adminCommonHandler{}
}

func (h *adminCommonHandler) Bind(router luchen.HTTPRouter) {
	router.Route("/admin/common", func(r chi.Router) {
		r.Handle("/insert", h.insert())
		r.Handle("/query", h.query())
		r.Handle("/get", h.get())
		r.Handle("/update", h.update())
		r.Handle("/delete", h.delete())
	})
}

func (h *adminCommonHandler) query() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeQueryEndpoint(),
		luchen.DecodeJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) get() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeGetEndpoint(),
		luchen.DecodeJSONRequest[daox.GetRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) insert() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeInsertEndpoint(),
		luchen.DecodeJSONRequest[daox.InsertRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) update() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeUpdateEndpoint(),
		luchen.DecodeJSONRequest[daox.UpdateRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

func (h *adminCommonHandler) delete() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeDeleteEndpoint(),
		luchen.DecodeJSONRequest[daox.DeleteRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

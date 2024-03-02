package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/transport/http"
)

type adminCommonHandler struct {
}

func newAdminCommonHandler() *adminCommonHandler {
	return &adminCommonHandler{}
}

func (h *adminCommonHandler) Bind(router *luchen.ServeMux) {
	commonAPI := http.AdminAPI + "/common"
	router.Handle(commonAPI+"/insert", h.insert())
	router.Handle(commonAPI+"/query", h.query())
	router.Handle(commonAPI+"/get", h.get())
	router.Handle(commonAPI+"/update", h.update())
	router.Handle(commonAPI+"/delete", h.delete())
}

func (h *adminCommonHandler) query() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(http.ErrorEncoder),
	}
	return luchen.NewHTTPHandler(
		MakeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
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
		luchen.DecodeHTTPJSONRequest[daox.GetRecord],
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
		luchen.DecodeHTTPJSONRequest[daox.InsertRecord],
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
		luchen.DecodeHTTPJSONRequest[daox.UpdateRecord],
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
		luchen.DecodeHTTPJSONRequest[daox.DeleteRecord],
		luchen.EncodeHTTPJSON(http.ResponseWrapper),
		options...,
	)
}

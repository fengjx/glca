package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fengjx/glca/protocol"
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
	router.Handle(commonAPI+"/batch-update", h.batchUpdate())
	router.Handle(commonAPI+"/delete", h.delete())
	router.Handle(commonAPI+"/delete-by-ids", h.deleteByIDs())
}

func (h *adminCommonHandler) query() *httptransport.Server {
	return http.NewHandler(
		MakeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) get() *httptransport.Server {
	return http.NewHandler(
		MakeGetEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.GetRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) insert() *httptransport.Server {
	return http.NewHandler(
		MakeInsertEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.InsertRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) update() *httptransport.Server {
	return http.NewHandler(
		MakeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.UpdateRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) batchUpdate() *httptransport.Server {
	return http.NewHandler(
		MakeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[protocol.BatchUpdateReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) delete() *httptransport.Server {
	return http.NewHandler(
		MakeDeleteEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.DeleteRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *adminCommonHandler) deleteByIDs() *httptransport.Server {
	return http.NewHandler(
		MakeDeleteByIDsEndpoint(),
		luchen.DecodeHTTPJSONRequest[protocol.DeleteByIDsReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/fengjx/go-halo/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/current"
)

const (
	StatusOK = 0

	AdminAPI = "/admin"
	OpenAPI  = "/api"

	RequestHeaderDebugUID = "X-Debug-UID"
	RequestHeaderToken    = "X-Token"

	ResponseHeaderServer       = "Server"
	ResponseHeaderRefreshToken = "X-Refresh-Token"
)

type result struct {
	Status int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func ResponseWrapper(data interface{}) interface{} {
	res := &result{
		Status: StatusOK,
		Msg:    "ok",
		Data:   data,
	}
	return res
}

// ErrorEncoder 统一异常处理
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log := luchen.Logger(ctx)
	var errn *luchen.Errno
	ok := errors.As(err, &errn)
	if !ok {
		log.Error("handler error", zap.Error(err))
		msg := errno.SystemErr.Msg
		if !luchen.IsProd() {
			msg = err.Error()
		}
		res := &result{
			Status: errno.SystemErr.Code,
			Msg:    msg,
		}
		WriteData(ctx, w, errno.SystemErr.HTTPCode, res)
		return
	}
	httpCode := 500
	if errn.HTTPCode > 0 {
		httpCode = errn.HTTPCode
	}
	res := &result{
		Status: errn.Code,
		Msg:    errn.Msg,
	}
	WriteData(ctx, w, httpCode, res)
}

func WriteData(ctx context.Context, w http.ResponseWriter, httpCode int, data any) {
	log := luchen.Logger(ctx)
	w.WriteHeader(httpCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error("write http response err", zap.Error(err))
	}
}

// NewHandler 创建 http handler
func NewHandler(e luchen.Endpoint,
	dec httptransport.DecodeRequestFunc,
	enc httptransport.EncodeResponseFunc,
	options ...httptransport.ServerOption) *httptransport.Server {

	options = append(options, httptransport.ServerErrorEncoder(ErrorEncoder))
	targetEndpoint := luchen.AccessMiddleware(&luchen.AccessLogOpt{
		PrintResp: true,
		ContextFields: map[string]luchen.GetValueFromContext{
			"uid": func(ctx context.Context) any {
				return current.UID(ctx)
			},
		},
	})(e)
	return luchen.NewHTTPHandler(
		targetEndpoint,
		dec,
		enc,
		options...,
	)
}

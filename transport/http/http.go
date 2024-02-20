package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/connom/errno"
)

const (
	pathAPI = "/openapi"
)

type result struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func httpResponseWrapper(data interface{}) interface{} {
	res := &result{
		Msg:  "ok",
		Data: data,
	}
	return res
}

// 统一返回值处理
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := &result{
		Msg:  "ok",
		Data: response,
	}
	logger := luchen.Logger(ctx)
	logger.Info("http response", zap.Any("data", res))
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

// 统一异常处理
func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log := luchen.Logger(ctx)

	writeData := func(res *result) {
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Error("write error msg fail", zap.Error(err))
		}
	}

	httpCode := 500
	var errn *luchen.Errno
	ok := errors.As(err, &errn)
	if !ok {
		log.Error("handler error", zap.Error(err))
		msg := errno.SystemErr.Msg
		if !luchen.IsProd() {
			msg = err.Error()
		}
		res := &result{
			Code: errno.SystemErr.Code,
			Msg:  msg,
		}
		writeData(res)
		return
	}
	if errn.HTTPCode > 0 {
		httpCode = errn.HTTPCode
	}
	w.WriteHeader(httpCode)
	res := &result{
		Code: errn.Code,
		Msg:  errn.Msg,
	}
	writeData(res)
}

package endpoint

import (
	"context"
	"errors"
	"path"

	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/fengjx/glca/current"
)

type Middleware = endpoint.Middleware

var accessLog = newLog()

func AccessMiddleware() Middleware {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			uid := current.UID(ctx)
			action := current.RequestAction(ctx)
			protocol := current.Protocol(ctx)
			method := current.Method(ctx)

			response, err = next(ctx, request)

			code := 0
			var errn *luchen.Errno
			ok := errors.As(err, &errn)
			if ok {
				code = errn.Code
			}
			accessLog.Info("",
				zap.Int64("uid", uid),
				zap.Any("request", request),
				zap.String("action", action),
				zap.String("protocol", protocol),
				zap.String("method", method),
				zap.Int("code", code),
				zap.String("err", err.Error()),
			)
			return
		}
	}
}

func newLog() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(luchen.GetLogDir(), "access.log"),
		MaxSize:    1024 * 1024 * 100,
		MaxBackups: 30,
		MaxAge:     30,
	})
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.FunctionKey = ""
	encoderConfig.LevelKey = ""
	encoderConfig.MessageKey = ""
	encoderConfig.NameKey = ""
	encoderConfig.CallerKey = ""

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zapcore.InfoLevel,
	)
	return zap.New(core, zap.AddCaller())
}

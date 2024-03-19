package endpoint

import (
	"context"

	"github.com/fengjx/luchen"
)

func Init(_ context.Context, httpServer *luchen.HTTPServer) {
	httpServer.Handler(newConfigHandler())
}

func MakeConfigFetchEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := luchen.Logger(ctx)
		_ = log
		return nil, nil
	}
}

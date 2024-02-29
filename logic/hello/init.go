package hello

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic/hello/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	endpoint.Init(ctx, httpServer)
}

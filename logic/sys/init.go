package sys

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic/sys/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	endpoint.Init(ctx, httpServer)
}

package common

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic/common/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	endpoint.Init(ctx, httpServer)
}

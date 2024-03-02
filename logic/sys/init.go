package sys

import (
	"context"

	"github.com/fengjx/luchen"

	_ "github.com/fengjx/glca/logic/sys/internal/dao"
	"github.com/fengjx/glca/logic/sys/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	endpoint.Init(ctx, httpServer)
}

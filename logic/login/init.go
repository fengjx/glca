package login

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic/login/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	endpoint.Init(ctx, httpServer)
}

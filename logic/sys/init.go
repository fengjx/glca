package sys

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/connom/lifecycle"
	"github.com/fengjx/glca/logic/sys/internal/dao"
	"github.com/fengjx/glca/logic/sys/internal/endpoint"
	"github.com/fengjx/glca/logic/sys/internal/provider"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	lifecycle.AddHook(lifecycle.InterfaceAware, func() {
		provider.Init()
	})
	lifecycle.AddHook(lifecycle.PostProcessor, func() {
		dao.RegisterTableConfig()
	})
	endpoint.Init(ctx, httpServer)
}

package sys

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/connom/lifecycle"
	"github.com/fengjx/glca/logic/sys/internal/dao"
	"github.com/fengjx/glca/logic/sys/internal/endpoint"
	"github.com/fengjx/glca/logic/sys/internal/provider"
	"github.com/fengjx/glca/logic/sys/syspub"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	lifecycle.AddHook(lifecycle.InterfaceAware, func() {
		syspub.SetUserAPI(provider.UserProvider)
	})
	lifecycle.AddHook(lifecycle.PostProcessor, func() {
		dao.RegisterTableConfig()
	})
	endpoint.Init(ctx, httpServer)
}

package common

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/connom/lifecycle"
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/common/internal/endpoint"
	"github.com/fengjx/glca/logic/common/internal/provider"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	lifecycle.AddHook(lifecycle.InterfaceAware, func() {
		commpub.SetCommonAPI(provider.CommonProvider)
	})
	endpoint.Init(ctx, httpServer)
}

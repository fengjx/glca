package logic

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic/common"
	"github.com/fengjx/glca/logic/hello"
	"github.com/fengjx/glca/logic/sys"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	hello.Init(ctx, httpServer)
	common.Init(ctx, httpServer)
	sys.Init(ctx, httpServer)
}
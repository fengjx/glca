package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/luchen"

	"github.com/fengjx/glca/logic"
	"github.com/fengjx/glca/transport/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := luchen.Logger(ctx)
	logger.Info("app start")

	httpServer := http.GetServer()
	logic.Init(ctx, httpServer)
	luchen.StartWithContext(ctx, httpServer)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	<-quit
	logger.Info("app stop")
	cancel()
	luchen.Stop()
}

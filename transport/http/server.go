package http

import (
	"sync"

	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/http/middleware"

	"github.com/fengjx/glca/connom/config"
)

var (
	server     *luchen.HTTPServer
	serverOnce sync.Once
)

func GetServer() *luchen.HTTPServer {
	serverOnce.Do(func() {
		serverConfig := config.GetConfig().Server.HTTP
		server = luchen.NewHTTPServer(
			luchen.WithServiceName(serverConfig.ServerName),
			luchen.WithServerAddr(serverConfig.Listen),
		).Use(
			middleware.Recoverer,
			middleware.RequestID,
			middleware.RealIP,
			middleware.AllowAll().Handler,
			adminAuthMiddleware,
		)
	})
	return server
}

package http

import (
	"net/http"
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
			middleware.CorsHandler(middleware.CorsOptions{
				AllowedOrigins: serverConfig.Cors.AllowOrigins,
				AllowedMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}),
			adminAuthMiddleware,
		).Static("/static/", "static")
	})
	return server
}

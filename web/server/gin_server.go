package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginServer struct {
	*gin.Engine
	RouterFunc func(*gin.Engine)
	Server     *http.Server
	Config     WebServerConfig
}

func newGinServer(config WebServerConfig) *ginServer {
	return &ginServer{
		Engine: gin.New(),
		Server: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
			TLSConfig:         nil,
			ReadTimeout:       config.ReadTimeout,
			ReadHeaderTimeout: config.ReadHeaderTimeout,
			WriteTimeout:      config.WriteTimeout,
			IdleTimeout:       0,
		},
		Config: config,
	}
}

// 外部注册进来的路由处理
func (gs *ginServer) initRouter() {
	if gs.RouterFunc != nil {
		gs.RouterFunc(gs.Engine)
	}
}

func (gs *ginServer) ListenAndServe() error {
	gs.initRouter()
	gs.Server.Handler = gs.Engine // 把gin路由交给net/http管理
	return gs.Server.ListenAndServe()
}

func (gs *ginServer) Shutdown(ctx context.Context) error {
	return gs.Server.Shutdown(ctx)
}

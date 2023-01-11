package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shixinshuiyou/framework/trace"
	"github.com/shixinshuiyou/framework/web/middleware"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	config       Config     // 配置
	mainServer   *ginServer // 应用服务
	statusServer *ginServer // 状态服务
	//signalTable  *signal.SignalTable // 信号表
	chServerStop chan struct{}  // 服务停止信号
	errGroup     errgroup.Group // 管理服务的启动
}

func NewServer(config Config) *Server {
	srv := &Server{
		config: config,
	}
	if config.Mode == ModeDebug {
		gin.SetMode(gin.DebugMode)
	}
	srv.mainServer = newGinServer(config.MainSrvConf)
	// 是否启用状态服务器
	if config.StatusSrvConf.Port > 0 {
		srv.statusServer = newGinServer(config.StatusSrvConf)
	}
	return srv
}

func (srv *Server) Start() {
	srv.initLog()
	srv.initRouter()
	srv.errGroup.Go(srv.mainServer.ListenAndServe)
	// TODO 后续切换成信号量控制
	srv.errGroup.Go(func() error {
		select {}
	})
	if srv.statusServer != nil {
		srv.statusServer.ListenAndServe()
	}
	if err := srv.errGroup.Wait(); err != nil {
		logrus.WithError(err).Panic("server init fault")
		return
	}
}

func (srv *Server) initLog() {
	// 注册Recovery:日志和方法
	srv.mainServer.Use(gin.RecoveryWithWriter(logrus.StandardLogger().Out, middleware.RecoveryMetric))
	// TODO trace日志
	if srv.config.TraceConf != trace.EmptyConfig {
		trace.NewJaegerTracer(srv.config.TraceConf)
		srv.mainServer.Use(otelgin.Middleware(srv.config.Name))
	}
}

func (srv *Server) initRouter() {
	if srv.config.CollectMetrics && srv.statusServer != nil {
		srv.mainServer.Use(middleware.MetricMiddleware(srv.config.Name))
		srv.RegisterMetric()
	}
	srv.mainServer.Use(middleware.SvcTagMiddleware)
	if srv.statusServer == nil {
		return
	}
	srv.statusServer.Use(middleware.SvcTagMiddleware)
}

// RegisterMetric 注册在状态服务器
func (srv *Server) RegisterMetric() {
	handler := promhttp.Handler()
	srv.statusServer.GET("/metrics", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})
}

// SetMainRouterFunc 添加路由钩子
func (srv *Server) SetMainRouterFunc(fn func(*gin.Engine)) {
	srv.mainServer.RouterFunc = fn
}

// SetStatusRouterFunc 添加状态服务钩子
func (srv *Server) SetStatusRouterFunc(fn func(*gin.Engine)) {
	srv.statusServer.RouterFunc = fn
}

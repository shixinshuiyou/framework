package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shixinshuiyou/framework/log"
	"github.com/shixinshuiyou/framework/signal"
	"github.com/shixinshuiyou/framework/trace"
	"github.com/shixinshuiyou/framework/web/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"golang.org/x/sync/errgroup"
	"os"
	"syscall"
	"time"
)

type Server struct {
	config       Config            // 配置
	mainServer   *ginServer        // 应用服务
	statusServer *ginServer        // 状态服务
	signalTable  *signal.SignTable // 信号表
	chServerStop chan struct{}     // 服务停止信号
	errGroup     errgroup.Group    // 管理服务的启动
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

	// default signal handler
	srv.signalTable = signal.NewSignTable()
	srv.signalTable.Register(syscall.SIGHUP, signal.IgnoreHandler)  // kill -1 is syscall.SIGHUP
	srv.signalTable.Register(syscall.SIGINT, srv.shutdownHandler)   // kill -2 is syscall.SIGINT
	srv.signalTable.Register(syscall.SIGQUIT, srv.shutdownHandler)  // kill -3 is syscall.SIGQUIT
	srv.signalTable.Register(syscall.SIGILL, signal.IgnoreHandler)  // kill -4 is syscall.SIGILL
	srv.signalTable.Register(syscall.SIGTRAP, signal.IgnoreHandler) // kill -5 is syscall.SIGTRAP
	srv.signalTable.Register(syscall.SIGABRT, signal.IgnoreHandler) // kill -6 is syscall.SIGABRT
	srv.signalTable.Register(syscall.SIGTERM, signal.TermHandler)   // kill -15 is syscall SIGTERM
	return srv
}

func (srv *Server) Start() {
	srv.initLog()
	srv.initRouter()
	srv.errGroup.Go(srv.mainServer.ListenAndServe)
	if srv.statusServer != nil {
		srv.statusServer.ListenAndServe()
	}
	srv.initSignTable()
	if err := srv.errGroup.Wait(); err != nil {
		log.Logger.Error("server init fault")
		return
	}

}

func (srv *Server) initLog() {
	log.InitLogger(srv.config.LogConf)
	// 注册Recovery:日志和方法
	srv.mainServer.Use(gin.RecoveryWithWriter(log.Logger.Writer(), middleware.RecoveryMetric))
	// gin-trace
	if srv.config.TraceConf != trace.EmptyConfig {
		_, err := trace.NewJaegerTracer(srv.config.TraceConf)
		if err != nil {
			log.Logger.Panic("jaeger tracer init fail")
		}
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

func (srv *Server) initSignTable() {
	srv.signalTable.StartSignalHandler()
}

// 关闭服务处理器
func (srv *Server) shutdownHandler(s os.Signal) {
	srv.Shutdown()
}

func (srv *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.mainServer.Shutdown(ctx); err != nil {
		log.Logger.Error("main Server Shutdown fail")
	}
	if srv.statusServer != nil {
		if err := srv.statusServer.Shutdown(ctx); err != nil {
			log.Logger.Error("Status Server Shutdown fail")
		}
	}
	<-ctx.Done()
	srv.signalTable.Shutdown()
	log.Logger.Info("server exiting")
}

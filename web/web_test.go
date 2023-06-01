package web

import (
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/database"
	"github.com/shixinshuiyou/framework/log"
	"github.com/shixinshuiyou/framework/trace"
	"github.com/shixinshuiyou/framework/web/server"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestStatusServer(t *testing.T) {
	config := server.Config{
		Mode:          server.ModeDebug,
		Name:          "czhTest",
		MainSrvConf:   server.WebServerConfig{Host: "0.0.0.0", Port: 10013},
		StatusSrvConf: server.WebServerConfig{Host: "0.0.0.0", Port: 10011},
		LogConf: log.Config{
			ServerTag: "test_local",
			Level:     logrus.InfoLevel,
		},
		TraceConf: trace.Config{
			Host:        "127.0.0.1",
			Port:        6831,
			ServiceName: "czhTest",
			Ratio:       1,
		},
		CollectMetrics: true,
	}
	nemServer := server.NewServer(config)
	dbConf := database.DatabaseConfig{
		Type:          "mysql",
		User:          "root",
		Password:      "chen360622",
		Host:          "127.0.0.1:3306",
		Name:          "mayo",
		MaxIdle:       2,
		MaxOpen:       10,
		LogLevel:      logrus.InfoLevel,
		SlowThreshold: 6 * time.Second,
		Colorful:      false,
	}
	db, _ := dbConf.InitDB()
	database.AddGormCallbacks(db, config.Name)

	nemServer.SetMainRouterFunc(func(engine *gin.Engine) {
		engine.GET("/test1", func(context *gin.Context) {
			var name string
			db.WithContext(context.Request.Context()).Table("mayo_user").Raw("select username from mayo_user where id = ?", 1).First(&name)
			context.JSONP(200, map[string]interface{}{
				"code": 0,
				"mes":  "hello " + name,
			})
		})
		engine.GET("/test2", func(context *gin.Context) {
			context.JSONP(200, map[string]interface{}{
				"code": 0,
				"mes":  "hello tt",
			})
		})
	})
	nemServer.SetStatusRouterFunc(func(engine *gin.Engine) {
		// 可以接入prometheus、或者热力图
	})

	nemServer.Start()

	//	resp, err := http.Get("http://127.0.0.1:10013/test1")
	//	resp, err := http.Get("http://127.0.0.1:10013/test2")
	//	httpResp, err := http.Get("http://127.0.0.1:10010/metrics")

}

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/web/server"
	"testing"
)

func TestStatusServer(t *testing.T) {
	config := server.Config{
		Mode:           server.ModeDebug,
		Name:           "web_test",
		MainSrvConf:    server.WebServerConfig{Host: "0.0.0.0", Port: 10013},
		StatusSrvConf:  server.WebServerConfig{Host: "0.0.0.0", Port: 10014},
		CollectMetrics: true,
	}
	server := server.NewServer(config)
	server.SetMainRouterFunc(func(engine *gin.Engine) {
		engine.GET("/test1", func(context *gin.Context) {
			context.JSONP(200, map[string]interface{}{
				"code": 0,
				"mes":  "hello yy",
			})
		})
		engine.GET("/test2", func(context *gin.Context) {
			context.JSONP(200, map[string]interface{}{
				"code": 0,
				"mes":  "hello tt",
			})
		})
	})

	server.Start()

	//	resp, err := http.Get("http://127.0.0.1:10013/test1")
	//	resp, err := http.Get("http://127.0.0.1:10013/test2")
	//	httpResp, err := http.Get("http://127.0.0.1:10014/metrics")

}

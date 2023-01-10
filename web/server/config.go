package server

import (
	"github.com/shixinshuiyou/framework/log"
	"github.com/shixinshuiyou/framework/trace"
	"time"
)

type Mode int

const (
	ModeDebug Mode = iota + 1
	ModeTest
	ModeProd
)

type Config struct {
	Mode           Mode            `json:"mode" yaml:"mode"`
	Name           string          `json:"name" yaml:"name"`
	MainSrvConf    WebServerConfig `json:"mainSerConf" yaml:"mainSerConf"`
	StatusSrvConf  WebServerConfig `json:"statusSerConf" yaml:"statusSerConf"`
	LogConf        log.Config      `json:"logConf" yaml:"logConf"`               // web 请求的访问日志
	TraceConf      trace.Config    `json:"traceConf" yaml:"traceConf"`           // 链路追踪
	CollectMetrics bool            `json:"collectMetrics" yaml:"collectMetrics"` // 是否开启指标采集
}

type WebServerConfig struct {
	Host              string        `json:"host" yaml:"host"`
	Port              int64         `json:"port" yaml:"port"`
	ReadHeaderTimeout time.Duration `json:"readHeaderTimeout" yaml:"readHeaderTimeout"` // ms, 不设置则不限制超时时间
	ReadTimeout       time.Duration `json:"read_timeout" yaml:"read_timeout"`           // ms, 不设置则不限制超时时间
	WriteTimeout      time.Duration `json:"write_timeout" yaml:"write_timeout"`         // ms, 不设置则不限制超时时间
	Templates         []string      `json:"templates" yaml:"templates"`                 // 模板目录
}

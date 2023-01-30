package log

import (
	"github.com/shixinshuiyou/framework/log/hook"
	"github.com/shixinshuiyou/framework/util/env"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func init() {
	defaultConf := Config{
		ServerTag: "dev",
		Level:     logrus.InfoLevel,
	}
	InitLogger(defaultConf)
}

func InitLogger(conf Config) *logrus.Entry {
	if conf.ServerTag == "" {
		conf.ServerTag = "dev"
	}
	logrus.SetLevel(conf.Level)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(true)
	Logger = logrus.WithField("env", env.GetEnv()).WithField("server_tag", conf.ServerTag)
	// 日志统一收集
	//addGraylogHook(conf.GrayLogConf)
	// 日志告警

	return Logger
}

func addGraylogHook(conf hook.GrayLogConfig) {
	grayHook := hook.NewGraylogHook(conf)
	logrus.AddHook(&grayHook)
}

func addAlarmHook(conf hook.AlarmHookConf) {
	alarmHook := hook.NewAlarmHook(conf)
	logrus.AddHook(&alarmHook)
}

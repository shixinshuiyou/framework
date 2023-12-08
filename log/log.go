package log

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/samber/lo"
	"github.com/shixinshuiyou/framework/log/hook"
	"github.com/shixinshuiyou/framework/version"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func InitLogger(conf Config) {
	if lo.IsEmpty(conf.AppFlag) {
		conf.AppFlag = "dev"
	}
	logrus.SetFormatter(&JSONFormatter{
		AppFlag:    conf.AppFlag,
		AppVersion: version.GetAppVersion(),
		GitCommit:  version.GetGitCommit(),
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(conf.Level)
	if lo.IsEmpty(conf.File) {
		logrus.SetOutput(os.Stdout)
		return
	}
	if conf.File.Size <= 0 {
		conf.File.Size = 100
	}
	if conf.File.Count <= 0 {
		conf.File.Count = 30
	}
	if lo.IsEmpty(conf.File.Path) {
		conf.File.Path = "log"
	}

	writer, err := rotatelogs.New(
		conf.AppFlag+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(conf.File.Path),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(uint(conf.File.Count)),
		rotatelogs.WithRotationSize(int64(conf.File.Size*1024*1024)),
	)

	fileAndStdoutWriter := io.MultiWriter(writer, os.Stdout)
	if err != nil {
		logrus.Fatal("initLog rotatelogs New err", err)
		return
	}
	logrus.SetOutput(fileAndStdoutWriter)

	// 日志统一收集
	//addGraylogHook(conf.GrayLogConf)
	// 日志告警
}

func addGraylogHook(conf hook.GrayLogConfig) {
	grayHook := hook.NewGraylogHook(conf)
	logrus.AddHook(&grayHook)
}

func addAlarmHook(conf hook.AlarmHookConf) {
	alarmHook := hook.NewAlarmHook(conf)
	logrus.AddHook(&alarmHook)
}

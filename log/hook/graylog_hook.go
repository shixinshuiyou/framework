package hook

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"strings"
)

type GrayLogConfig struct {
	Proto string // 传输方式
	Addr  string
	Level logrus.Level
}

type GraylogHook struct {
	Level  logrus.Level
	writer gelf.Writer
}

func NewGraylogHook(conf GrayLogConfig) (hook GraylogHook) {
	var err error
	switch strings.ToLower(conf.Proto) {
	case "udp":
		hook.writer, err = gelf.NewUDPWriter(conf.Addr)
	case "tcp":
	default:
		hook.writer, err = gelf.NewTCPWriter(conf.Addr)
	}
	if err != nil {
		logrus.Error(err)
	}
	hook.Level = conf.Level
	return
}

func (hook *GraylogHook) Fire(entry *logrus.Entry) error {
	_, err := hook.writer.Write([]byte(entry.Message))
	logrus.Error(err)
	return err
}

func (hook *GraylogHook) Levels() []logrus.Level {
	var levels []logrus.Level
	for _, level := range logrus.AllLevels {
		if level <= hook.Level {
			levels = append(levels, level)
		}
	}
	return levels
}

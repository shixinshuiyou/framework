package log

import (
	"github.com/shixinshuiyou/framework/log/hook"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerTag   string              // 唯一标识（平台-项目-环境）
	Level       logrus.Level        // 日志级别
	GrayLogConf *hook.GrayLogConfig // 统一收集日志配置
}

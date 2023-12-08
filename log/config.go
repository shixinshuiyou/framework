package log

import (
	"github.com/shixinshuiyou/framework/log/hook"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppFlag     string              // 唯一标识（平台-项目-环境）
	Level       logrus.Level        // 日志级别
	GrayLogConf *hook.GrayLogConfig // 统一收集日志配置
	File        *FileConfig         // 本地持久化方案
}

type FileConfig struct {
	Path  string // 文件持久化路经
	Size  int    // 单个文件大小（单位：M）
	Count int    // 文件数量
}

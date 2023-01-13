package hook

import "github.com/sirupsen/logrus"

type AlarmHook struct{}

type AlarmHookConf struct {
}

func NewAlarmHook(conf AlarmHookConf) AlarmHook {
	return AlarmHook{}
}

func (hook *AlarmHook) Fire(entry *logrus.Entry) error {
	// TODO 实现日志告警代码
	return nil
}

func (hook *AlarmHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

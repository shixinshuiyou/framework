package log

import (
	"fmt"
	_const "github.com/shixinshuiyou/framework/const"
	"github.com/shixinshuiyou/framework/util"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	AppFlag    string
	AppVersion string
	GitCommit  string
	logrus.JSONFormatter
}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["app_name"] = f.AppFlag
	entry.Data["app_version"] = f.AppVersion
	entry.Data["git_commit"] = f.GitCommit
	if ctx := entry.Context; ctx != nil {
		if traceID := util.GetContextValue(ctx, _const.TraceID); traceID != "" {
			entry.Data["trace_id"] = traceID
		}
		if accountID := util.GetContextValue(ctx, _const.UserID); accountID != "" {
			entry.Data["account_id"] = accountID
		}
		if accountName := util.GetContextValue(ctx, _const.UserName); accountName != "" {
			entry.Data["account_name"] = accountName
		}
	}

	f.JSONFormatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		filePath := fmt.Sprintf("%s/%s:%d", path.Dir(frame.Function), path.Base(frame.File), frame.Line)
		return path.Base(frame.Function), filePath
	}
	return f.JSONFormatter.Format(entry)
}

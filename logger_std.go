package wkhtml

import (
	"fmt"
	"runtime"
	"strings"
)

// loggerNil 标准日志器。
type loggerStd struct {
	logLevelToTagMap map[int32]string
}

// NewLoggerStd 新建标准日志器。
func NewLoggerStd() *loggerStd {
	return &loggerStd{
		logLevelToTagMap: map[int32]string{
			logLevelDebug: "debug",
			logLevelInfo:  "info",
			logLevelWarn:  "warn",
			logLevelError: "error",
		},
	}
}

// Log 输出日志。
func (l *loggerStd) Log(logLevel int32, logInfo string, fields ...*field) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		if len(fields) != 0 {
			fieldInfos := make([]string, 0)
			for _, field := range fields {
				fieldInfos = append(fieldInfos, fmt.Sprintf("{%s: %s}", field.key, field.value))
			}
			fmt.Printf("[%s] %s:%d %s | %s\n", l.getLogTag(logLevel), file, line, logInfo,
				strings.Join(fieldInfos, ","))
			return
		}
		fmt.Printf("[%s] %s:%d %s\n", l.getLogTag(logLevel), file, line, logInfo)
	}
}

// getLogTag 获取日志标志。
func (l *loggerStd) getLogTag(logLevel int32) string {
	logTag, ok := l.logLevelToTagMap[logLevel]
	if ok {
		return logTag
	}
	return "unknown"
}

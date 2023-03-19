package wkhtml

// loggerNil 空日志器。
type loggerNil struct{}

// NewLoggerNil 新建空日志器。
func NewLoggerNil() *loggerNil {
	return &loggerNil{}
}

// Log 输出日志。
func (l *loggerNil) Log(logLevel int32, logInfo string, fields ...*field) {}

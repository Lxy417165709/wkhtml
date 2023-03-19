package wkhtml

// init 初始化函数。
func init() {
	globalLogger = NewLoggerNil()
}

// globalLogger 全局日志器。
var globalLogger Logger

// SetLogger 设置日志器。
func SetLogger(logger Logger) {
	globalLogger = logger
}

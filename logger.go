package wkhtml

// Logger 日志器。
type Logger interface {
	Log(logLevel int32, logInfo string, logFields ...*field)
}

// field 字段。
type field struct {
	key   string
	value string
}

// Field 新建字段。
func Field(key, value string) *field {
	return &field{
		key:   key,
		value: value,
	}
}

// 日志级别。
const (
	logLevelDebug int32 = 1
	logLevelInfo  int32 = 2
	logLevelWarn  int32 = 3
	logLevelError int32 = 4
)

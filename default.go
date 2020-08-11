package vlogger

// 默认日志输出
var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

// Reset will remove all the adapter
func Reset() {
	defaultLogger.Reset()
}

// SetTrimPath -
func SetTrimPath(trimPath string) {
	defaultLogger.SetTrimPath(trimPath)
}

// Painc logs a message at emergency level and panic.
func Painc(f interface{}, v ...interface{}) {
	defaultLogger.Panic(formatLog(f, v...))
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	defaultLogger.Error(formatLog(f, v...))
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	defaultLogger.Debug(formatLog(f, v...))
}

// Success logs a message at info level.
func Success(f interface{}, v ...interface{}) {
	defaultLogger.Success(formatLog(f, v...))
}

// Failed logs a message at info level.
func Failed(f interface{}, v ...interface{}) {
	defaultLogger.Failed(formatLog(f, v...))
}

// Normal logs a message at info level.
func Normal(f interface{}, v ...interface{}) {
	defaultLogger.Normal(formatLog(f, v...))
}

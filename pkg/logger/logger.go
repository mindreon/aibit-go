package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger 结构化日志实例
type Logger struct {
	zerolog.Logger
}

// 全局日志实例
var globalLogger *Logger

// Init 初始化日志
func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	globalLogger = &Logger{Logger: logger}
}

// GetLogger 获取全局日志实例
func GetLogger() *Logger {
	if globalLogger == nil {
		Init()
	}
	return globalLogger
}

// Info 信息日志
func (l *Logger) Info(msg string, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Info().Msg(msg)
		return
	}
	event := l.Logger.Info()
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			event = event.Interface(args[i].(string), args[i+1])
		}
	}
	event.Msg(msg)
}

// Error 错误日志
func (l *Logger) Error(msg string, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Error().Msg(msg)
		return
	}
	event := l.Logger.Error()
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			event = event.Interface(args[i].(string), args[i+1])
		}
	}
	event.Msg(msg)
}

// Debug 调试日志
func (l *Logger) Debug(msg string, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Debug().Msg(msg)
		return
	}
	event := l.Logger.Debug()
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			event = event.Interface(args[i].(string), args[i+1])
		}
	}
	event.Msg(msg)
}

// Warn 警告日志
func (l *Logger) Warn(msg string, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Warn().Msg(msg)
		return
	}
	event := l.Logger.Warn()
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			event = event.Interface(args[i].(string), args[i+1])
		}
	}
	event.Msg(msg)
}

// 便民函数
func Info(msg string, args ...interface{}) {
	GetLogger().Info(msg, args...)
}

func Error(msg string, args ...interface{}) {
	GetLogger().Error(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	GetLogger().Debug(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	GetLogger().Warn(msg, args...)
}

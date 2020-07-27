package logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type API interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close() error
}

type metaData struct {
	time       time.Time
	level      Level
	fileName   string
	funcName   string
	lineNumber int
	content    string
}

type Logger struct {
	file    API
	console API
	opts    options
	context context.Context
	cancel  context.CancelFunc
}

func New(opts ...Option) *Logger {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	file := NewFile(o.level, o.path, ctx)
	console := NewConsole(o.level)
	return &Logger{
		opts:    o,
		file:    file,
		console: console,
		context: ctx,
		cancel:  cancelFunc,
	}
}

func levelText(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "INFO"
	}
}

// 设置日志级别
func (logger *Logger) SetLevel(level Level) {
	if level <= DebugLevel || level > FatalLevel {
		level = DebugLevel
	}
	logger.opts.level = level
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	logger.file.Debug(format, args...)
	logger.console.Debug(format, args...)
}

func (logger *Logger) Trace(format string, args ...interface{}) {
	logger.file.Trace(format, args...)
	logger.console.Trace(format, args...)
}

func (logger *Logger) Info(format string, args ...interface{}) {
	logger.file.Info(format, args...)
	logger.console.Info(format, args...)
}

func (logger *Logger) Warn(format string, args ...interface{}) {
	logger.file.Warn(format, args...)
	logger.console.Warn(format, args...)
}

func (logger *Logger) Error(format string, args ...interface{}) {
	logger.file.Error(format, args...)
	logger.console.Error(format, args...)
}

func (logger *Logger) Fatal(format string, args ...interface{}) {
	logger.file.Fatal(format, args...)
	logger.console.Fatal(format, args...)
}

func (logger *Logger) Close() {
	logger.cancel()
	logger.file.Close()
	logger.console.Close()
}

// PathExists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// GetLineInfo
func GetLineInfo() (fileName, funcName string, lineNumber int) {
	pc, file, line, ok := runtime.Caller(5)
	if ok {
		fileName = path.Base(file)
		funcName = runtime.FuncForPC(pc).Name()
		lineNumber = line
	}
	return
}

func fields(level Level, format string, args ...interface{}) *metaData {
	fileName, funcName, lineNumber := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	content := fmt.Sprintf(format, args...)

	return &metaData{
		time:       time.Now(),
		level:      level,
		fileName:   fileName,
		funcName:   funcName,
		lineNumber: lineNumber,
		content:    content,
	}
}

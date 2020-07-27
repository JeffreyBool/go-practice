package logger

import (
	"fmt"
	"os"
)

type Console struct {
	level Level
}

func NewConsole(level Level) API {
	return &Console{
		level: level,
	}
}

func (c *Console) write(fd *os.File, level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	meta := fields(level, format, args...)
	fmt.Fprintf(fd, "%s %s (%s:%s:%d) %s\n", meta.time, levelText(meta.level), meta.fileName, meta.funcName, meta.lineNumber, meta.content)
}

func (c Console) Debug(format string, args ...interface{}) {
	c.write(os.Stdout, DebugLevel, format, args...)
}

func (c Console) Trace(format string, args ...interface{}) {
	c.write(os.Stdout, TraceLevel, format, args...)
}

func (c Console) Info(format string, args ...interface{}) {
	c.write(os.Stdout, InfoLevel, format, args...)
}

func (c Console) Warn(format string, args ...interface{}) {
	c.write(os.Stdout, WarnLevel, format, args...)
}

func (c Console) Error(format string, args ...interface{}) {
	c.write(os.Stderr, ErrorLevel, format, args...)
}

func (c Console) Fatal(format string, args ...interface{}) {
	c.write(os.Stderr, FatalLevel, format, args...)
}

func (c Console) Close() error {
	return nil
}

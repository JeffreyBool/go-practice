package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type File struct {
	level  Level
	path   string
	name   string
	fd     *os.File
	warnFd *os.File
}

func NewFile(level Level, path string, name string) Logger {
	file := &File{level: level, path: path, name: name}
	file.init()
	return file
}

func (f *File) init() {
	if !PathExists(f.path) {
		if err := os.MkdirAll(f.path, os.ModePerm); err != nil {
			panic(err)
		}
	}

	filename := fmt.Sprintf("%s/%s.log", f.path, f.name)
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(errors.Errorf("open file %s failed, err: %v", filename, err))
	}
	f.fd = fd

	filename = fmt.Sprintf("%s/%s.error.log", f.path, f.name)
	fd, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(errors.Errorf("open file %s failed, err: %v", filename, err))
	}
	f.warnFd = fd
}

func (f *File) SetLevel(level Level) {
	if level <= DebugLevel || level > FatalLevel {
		level = DebugLevel
	}
	f.level = level
}

func (f *File) Debug(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.fd)
	fmt.Fprint(f.fd, msg)
}

func (f *File) Trace(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.fd)
	fmt.Fprint(f.fd, msg)
}

func (f *File) Info(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.fd)
	fmt.Fprint(f.fd, msg)
}

func (f File) Warn(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.fd)
	fmt.Fprint(f.fd, msg)
}

func (f File) Error(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.warnFd)
	fmt.Fprint(f.warnFd, msg)
}

func (f File) Fatal(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(f.level)
	msg := nowStr + " " + levelStr + " " + fmt.Sprintf(format, args...)
	fmt.Fprintln(f.warnFd)
	fmt.Fprint(f.warnFd, msg)
}

func (f File) Close() error {
	f.fd.Close()
	f.warnFd.Close()
	return nil
}

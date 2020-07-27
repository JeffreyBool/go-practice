package logger

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type File struct {
	level   Level
	path    string
	fd      *os.File
	errorFd *os.File
	queue   chan *metaData
	context context.Context
}

func NewFile(level Level, path string, context context.Context) API {
	file := &File{
		level:   level,
		path:    path,
		queue:   make(chan *metaData, 20000),
		context: context,
	}
	file.init()
	go file.loopWrite()
	return file
}

func (f *File) loopWrite() {
	infoFdWrite := bufio.NewWriter(f.fd)
	errFdWrite := bufio.NewWriter(f.errorFd)
	for meta := range f.queue {
		select {
		case <-f.context.Done():
			goto Stop
		default:
		}

		msg := fmt.Sprintf("%s %s (%s:%s:%d) %s\n", meta.time, levelText(meta.level), meta.fileName, meta.funcName, meta.lineNumber, meta.content)
		if meta.level >= ErrorLevel {
			if _, err := errFdWrite.WriteString(msg); err != nil {
				os.Stdout.WriteString(err.Error())
			}
		} else {
			if _, err := infoFdWrite.WriteString(msg); err != nil {
				os.Stdout.WriteString(err.Error())
			}
		}
		infoFdWrite.Flush()
		errFdWrite.Flush()
	}
Stop:
}

func (f *File) init() {
	if !PathExists(path.Dir(f.path)) {
		if err := os.MkdirAll(path.Dir(f.path), os.ModePerm); err != nil {
			panic(err)
		}
	}

	filename := f.path
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(errors.Errorf("open file %s failed, err: %v", filename, err))
	}
	f.fd = fd

	var pl []string
	if pl = strings.Split(path.Base(f.path), ".log"); len(pl) == 0 {
		panic("请输入文件名称")
	}
	filename = fmt.Sprintf("%s/%s.error.log", path.Dir(f.path), pl[0])
	fd, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(errors.Errorf("open file %s failed, err: %v", filename, err))
	}
	f.errorFd = fd
}

func (f *File) write(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}

	data := fields(level, format, args...)
	select {
	case f.queue <- data:
	case <-f.context.Done():
	default:
	}
}

func (f *File) Debug(format string, args ...interface{}) {
	f.write(DebugLevel, format, args...)
}

func (f *File) Trace(format string, args ...interface{}) {
	f.write(TraceLevel, format, args...)
}

func (f *File) Info(format string, args ...interface{}) {
	f.write(InfoLevel, format, args...)
}

func (f File) Warn(format string, args ...interface{}) {
	f.write(WarnLevel, format, args...)
}

func (f File) Error(format string, args ...interface{}) {
	f.write(ErrorLevel, format, args...)
}

func (f File) Fatal(format string, args ...interface{}) {
	f.write(FatalLevel, format, args...)
}

func (f File) Close() error {
	f.fd.Close()
	f.errorFd.Close()
	return nil
}

/**
* Author: JeffreyBool
* Date: 2020/7/27
* Time: 16:02
* Software: GoLand
 */

package logger

type Level int

const (
	DebugLevel Level = iota + 1
	TraceLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type options struct {
	level Level
	path  string
}

type Option func(*options)

// SetLevel 设置级别
func SetLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// SetPath 设置路径
func SetPath(path string) Option {
	return func(o *options) {
		o.path = path
	}
}

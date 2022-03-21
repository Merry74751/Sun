package logger

import (
	"github.com/Merry74751/sum/util"
	"log"
	"runtime/debug"
)

var sumLog *sum

func Logger() SumLogger {
	return sumLog
}

type SumLogger interface {
	Error(src string, params ...any)
	Warning(src string, params ...any)
	Info(src string, params ...any)
	Debug(src string, params ...any)
	Enable
}

type Enable interface {
	IsEnableError() bool
	IsEnableWarning() bool
	IsEnableInfo() bool
	IsEnableDebug() bool
}

type sum struct {
	error   *log.Logger
	warning *log.Logger
	info    *log.Logger
	debug   *log.Logger
}

func (s sum) Error(src string, params ...any) {
	if s.IsEnableError() {
		src = util.Formatter(src, params...)
		src = util.ColorText(util.Red, src)
		bytes := debug.Stack()
		src = src + "\n" + string(bytes)
		s.error.Println(src)
	}
}

func (s sum) Warning(src string, params ...any) {
	if s.IsEnableWarning() {
		src = util.Formatter(src, params...)
		src = util.ColorText(util.Cyan, src)
		s.warning.Println(src)
	}
}

func (s sum) Info(src string, params ...any) {
	if s.IsEnableInfo() {
		src = util.Formatter(src, params...)
		s.info.Println(src)
	}
}

func (s sum) Debug(src string, params ...any) {
	if s.IsEnableDebug() {
		src = util.Formatter(src, params...)
		s.debug.Println(src)
	}
}

func (s sum) IsEnableError() bool {
	return config.level >= DEBUG && config.level <= ERROR
}

func (s sum) IsEnableWarning() bool {
	return config.level >= DEBUG && config.level <= WARNING
}

func (s sum) IsEnableInfo() bool {
	return config.level >= DEBUG && config.level <= INFO
}

func (s sum) IsEnableDebug() bool {
	return config.level == DEBUG
}

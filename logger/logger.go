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
	isEnableError() bool
	isEnableWarning() bool
	isEnableInfo() bool
	isEnableDebug() bool
}

type sum struct {
	error   *log.Logger
	warning *log.Logger
	info    *log.Logger
	debug   *log.Logger
}

func (s sum) Error(src string, params ...any) {
	if s.isEnableError() {
		src = util.Formatter(src, params...)
		src = util.ColorText(util.Red, src)
		bytes := debug.Stack()
		src = src + "\n" + string(bytes)
		s.error.Println(src)
	}
	if enableWriteLog() {
		if config.writeFileLevel >= DEBUG && config.writeFileLevel <= ERROR {
			go WriteLog("error", src)
		}
	}
}

func (s sum) Warning(src string, params ...any) {
	if s.isEnableWarning() {
		src = util.Formatter(src, params...)
		src = util.ColorText(util.Cyan, src)
		s.warning.Println(src)
	}
	if enableWriteLog() {
		if config.writeFileLevel >= DEBUG && config.writeFileLevel <= WARNING {
			go WriteLog("warning", src)
		}
	}
}

func (s sum) Info(src string, params ...any) {
	if s.isEnableInfo() {
		src = util.Formatter(src, params...)
		s.info.Println(src)
	}
	if enableWriteLog() {
		if config.writeFileLevel >= DEBUG && config.writeFileLevel <= INFO {
			go WriteLog("info", src)
		}
	}
}

func (s sum) Debug(src string, params ...any) {
	if s.isEnableDebug() {
		src = util.Formatter(src, params...)
		s.debug.Println(src)
	}
	if enableWriteLog() {
		if config.writeFileLevel == DEBUG {
			go WriteLog("debug", src)
		}
	}
}

func (s sum) isEnableError() bool {
	return config.level >= DEBUG && config.level <= ERROR
}

func (s sum) isEnableWarning() bool {
	return config.level >= DEBUG && config.level <= WARNING
}

func (s sum) isEnableInfo() bool {
	return config.level >= DEBUG && config.level <= INFO
}

func (s sum) isEnableDebug() bool {
	return config.level == DEBUG
}

func enableWriteLog() bool {
	return config.enableWriteFile
}

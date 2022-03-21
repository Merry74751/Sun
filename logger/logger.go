package logger

import (
	"github.com/Merry74751/sum/util"
	"log"
	"runtime/debug"
)

var sumLog *sum

type SumLogger interface {
	Error(src string, params ...any)
	Warning(src string, params ...any)
	Info(src string, params ...any)
	Debug(src string, params ...any)
}

type sum struct {
	error   *log.Logger
	warning *log.Logger
	info    *log.Logger
	debug   *log.Logger
}

func (s sum) Error(src string, params ...any) {
	src = util.Formatter(src, params...)
	src = util.ColorText(util.Red, src)
	bytes := debug.Stack()
	src = src + "\n" + string(bytes)
	s.error.Println(src)
}

func (s sum) Warning(src string, params ...any) {
	src = util.Formatter(src, params...)
	src = util.ColorText(util.Cyan, src)
	s.warning.Println(src)
}

func (s sum) Info(src string, params ...any) {
	src = util.Formatter(src, params...)
	s.info.Println(src)
}

func (s sum) Debug(src string, params ...any) {
	src = util.Formatter(src, params...)
	s.debug.Println(src)
}

func Error(src string, params ...any) {
	sumLog.Error(src, params...)
}

func Warning(src string, params ...any) {
	sumLog.Warning(src, params...)
}

func Info(src string, params ...any) {
	sumLog.Info(src, params...)
}

func Debug(src string, params ...any) {
	sumLog.Debug(src, params...)
}

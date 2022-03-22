package logger

import (
	"github.com/Merry74751/sum/util"
	"os"
	"strings"
	"time"
)

const (
	FileSuffix    = ".log"
	PathSeparated = string(os.PathSeparator)
)

func WriteLog(level, content string) {
	path := getPath(config.path)
	fileName := getFileName(level, config.fileName)
	util.Mkdir(path)
	file := path + PathSeparated + fileName
	util.WriteFile(file, content)
}

func getFileName(level, filename string) string {
	date := time.Now().Format("2006-01-02")
	builder := strings.Builder{}
	builder.WriteString(filename)
	builder.WriteString(".")
	builder.WriteString(level)
	builder.WriteString(".")
	builder.WriteString(date)
	builder.WriteString(FileSuffix)
	return builder.String()
}

func getPath(path string) string {
	path = path + PathSeparated + util.GetDateDri()
	return path
}

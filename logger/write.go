package logger

import (
	"os"
)

const (
	FileSuffix    = ".log"
	PathSeparated = string(os.PathSeparator)
)

func writeFile(level, content string) {
	path := config.path
	checkPath(path)

	fileName := config.fileName + "-" + level
	file := path + PathSeparated + fileName + FileSuffix

}

func checkPath(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(path, os.ModeDir)
			if err != nil {
				sumLog.Warning("mkdir folder: {} error {}", path, err)
			}
		}
	}
}

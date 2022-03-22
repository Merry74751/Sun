package util

import (
	"log"
	"os"
	"strings"
	"time"
)

func Mkdir(dirname string) {
	if !DirIsExist(dirname) {
		err := os.MkdirAll(dirname, 0777)
		if err != nil {
			log.Printf("mkdir %s error: %s", dirname, err)
		}
	}
}

func DirIsExist(dirname string) bool {
	_, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetDateDri() string {
	format := time.Now().Format("2006/01/02")
	format = strings.Replace(format, "/", string(os.PathSeparator), 2)
	return format
}

func WriteFile(file string, c string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("openFile error: %s", err)
	}
	_, err1 := f.WriteString(c + "\n")
	if err1 != nil {
		log.Printf("writeFile error: %s", err)
	}
	f.Close()
}

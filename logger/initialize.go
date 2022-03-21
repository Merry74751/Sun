package logger

import (
	"github.com/Merry74751/sum/util"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	initConfig()
	initSumLog()
}

func initSumLog() {
	sumLog = new(sum)
	sumLog.error = log.New(os.Stderr, util.ColorText(util.Red, "Error   "), log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.warning = log.New(os.Stderr, util.ColorText(util.Cyan, "Warning "), log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.info = log.New(os.Stderr, "Info    ", log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.debug = log.New(os.Stderr, "Debug   ", log.Ldate|log.Ltime|log.Lshortfile)
}

func initConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Println("failed to get current path, error: {}", err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("failed to read the configuration file, error: {}", err)
	}

	config = new(sumConfig)
	level := viper.Get("logger.level")
	if level == nil {
		config.level = INFO
	} else {
		v := level.(string)
		config.level = getLevel(v)
	}
}

func getLevel(str string) int {
	switch str {
	case "info":
		return INFO
	case "debug":
		return DEBUG
	case "error":
		return ERROR
	case "warning":
		return WARNING
	default:
		log.Printf("config.yaml logger.level value was error: %s, value must be 'error', 'warning', 'info', 'debug'", str)
	}
	return -1
}

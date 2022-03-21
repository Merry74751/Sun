package logger

import (
	"github.com/Merry74751/sum/util"
	"github.com/spf13/viper"
	"log"
	"os"
)

var vip *viper.Viper = new(viper.Viper)

func init() {
	initSumLog()
	initConfig()
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
		Info("failed to get current path, error: {}", err)
	}
	vip.SetConfigName("config")
	vip.AddConfigPath(path)
	vip.SetConfigType("yaml")

	err = vip.ReadInConfig()
	if err != nil {
		Info("failed to read the configuration file, error: {}", err)
	}

	config = new(sumConfig)
	level := vip.Get("logger.level")
	if level == nil {
		config.level = "info"
	} else {
		config.level = level.(string)
	}

}

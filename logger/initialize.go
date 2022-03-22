package logger

import (
	"github.com/Merry74751/sum/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func init() {
	initConfig()
	initSumLog()
	sumLog.Info("Reading the configuration file is complete")
}

func initSumLog() {
	sumLog = new(sum)
	sumLog.error = log.New(os.Stderr, util.ColorText(util.Red, "Error   "), log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.warning = log.New(os.Stderr, util.ColorText(util.Cyan, "Warning "), log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.info = log.New(os.Stderr, "Info    ", log.Ldate|log.Ltime|log.Lshortfile)
	sumLog.debug = log.New(os.Stderr, "Debug   ", log.Ldate|log.Ltime|log.Lshortfile)
}

func initConfig() {
	projectPath, err := os.Getwd()
	if err != nil {
		log.Println("failed to get current path, error: {}", err)
		return
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(projectPath)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("failed to read the configuration file, error: {}", err)
		return
	}

	config = new(sumConfig)
	level := viper.GetString("logger.level")
	if level == "" {
		config.level = INFO
	} else {
		config.level = getLevel(level)
	}

	enableWriteFile := viper.GetBool("logger.enableWriteFile")
	config.enableWriteFile = enableWriteFile
	if !enableWriteFile {
		return
	}

	path := viper.GetString("logger.path")
	if path == "" {
		err := util.ColorText(util.Red, "config.yaml logger.path value is nil!")
		log.Println(err)
		return
	}
	config.path = path

	filename := viper.GetString("logger.filename")
	if filename == "" {
		err := util.ColorText(util.Red, "config.yaml logger.filename value is nil")
		log.Println(err)
		return
	}
	config.fileName = filename

	writeFileLevel := viper.GetString("logger.writeFileLevel")
	if writeFileLevel == "" {
		err := util.ColorText(util.Red, "config.yaml logger.writeFileLevel value is nil")
		log.Println(err)
		return
	}
	config.writeFileLevel = getLevel(writeFileLevel)
}

func getLevel(str string) int {
	str = strings.ToLower(str)
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

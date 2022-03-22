package Sun

import (
	"github.com/Merry74751/sum/logger"
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	sumLog := logger.Logger()
	sumLog.Info("test:{}", "test")
	sumLog.Debug("test:{}", "test")
}

func TestWriteFile(t *testing.T) {
	logger.WriteLog("info", "test")
}

func TestWriteFile2(t *testing.T) {
	path := "D:\\home"
	filename := path + string(os.PathSeparator) + "test.log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		t.Log("创建失败", err)
	}
	_, err1 := file.WriteString("do test\n")
	if err1 != nil {
		t.Log(err1)
	}
	file.Close()
}

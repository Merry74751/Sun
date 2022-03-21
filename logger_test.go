package Sun

import (
	"github.com/Merry74751/sum/logger"
	"testing"
)

func TestInfo(t *testing.T) {
	sumLog := logger.Logger()
	sumLog.Info("test:{}", "test")
	sumLog.Debug("test:{}", "test")
}

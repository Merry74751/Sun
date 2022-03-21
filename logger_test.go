package Sun

import (
	"github.com/Merry74751/sum/logger"
	"testing"
)

func TestInfo(t *testing.T) {
	logger.Error("test: {}", "test")
}

package interfaces_test

import (
	"testing"

	"github.com/ireina7/fgo/interfaces"
)

func TestLogging(t *testing.T) {
	logger := interfaces.NewPreludeLogger(nil)
	logger.Debug("This is a debug: %v", 7)
	logger.Info("This is a info: %v", 0)
	logger.Warn("This is a warning: %v", 0)
}

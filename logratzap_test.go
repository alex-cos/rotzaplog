package rotzaplog_test

import (
	"testing"

	"github.com/alex-cos/rotzaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestConsole(t *testing.T) {
	t.Parallel()

	logger := rotzaplog.InitConsoleLogger(zapcore.DebugLevel, true)
	defer logger.Sync() // nolint: errcheck
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Info("Test")
}

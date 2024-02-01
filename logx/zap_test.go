package logx

import (
	"github.com/wordpress-plus/kit-logger/logx/config"
	"testing"
)

func TestZap(t *testing.T) {
	zap := config.Zap{
		Level:         "info",
		Format:        "console",
		Prefix:        "[github.com/wordpress-plus/server-core]",
		Director:      "log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}

	LOGGER := Zap(zap)
	LOGGER.Info("test log output")
}

package zapx

import (
	"github.com/wordpress-plus/kit-logger/zapx/config"
	"testing"
)

func TestZap(t *testing.T) {
	zap := config.Zap{
		Level:         "info",
		Format:        "json",
		Prefix:        "[kit-logger]",
		Director:      "log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}

	LOGGER := Zap(zap)
	LOGGER.Info("test log output")
}

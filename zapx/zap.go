package zapx

import (
	"fmt"
	"github.com/wordpress-plus/kit-logger/zapx/config"
	"github.com/wordpress-plus/kit-logger/zapx/internal"
	"github.com/wordpress-plus/kit-logger/zapx/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zapx.Logger
func Zap(config config.Zap) (logger *zap.Logger) {
	if ok, _ := util.PathExists(config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Director)
		_ = os.Mkdir(config.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores(config)
	logger = zap.New(zapcore.NewTee(cores...))

	if config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func NewZap() (logger *zap.Logger) {
	return Zap(config.Zap{
		Format: "json",
	})
}

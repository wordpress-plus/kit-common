package logx

import (
	"fmt"
	"github.com/wordpress-plus/kit-logger/logx/config"
	"github.com/wordpress-plus/kit-logger/logx/internal"
	"github.com/wordpress-plus/kit-logger/logx/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
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

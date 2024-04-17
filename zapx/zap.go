package zapx

import (
	"fmt"
	kg "github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/zapx/config"
	"github.com/wordpress-plus/kit-common/zapx/internal"
	"github.com/wordpress-plus/kit-common/zapx/util"
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

func InitZap() *zap.Logger {
	kg.L = Zap(kg.C.Zap)
	zap.ReplaceGlobals(kg.L)
	return kg.L
}

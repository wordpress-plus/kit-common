package viperx

import (
	"github.com/spf13/viper"
	kg "github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/viperx/initialize"
)

func InitViper(path ...string) {
	// init viper
	kg.V = initialize.Viper(&kg.C, path...) // 初始化Viper
}

func InitViperV2(conf any, path ...string) *viper.Viper {
	// init viper
	kg.V = initialize.Viper(&kg.C, path...) // 初始化Viper
	initialize.Viper(conf, path...)         // 初始化Viper

	return kg.V
}

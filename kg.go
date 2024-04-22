package kit

import (
	"github.com/micro-services-roadmap/kit-common/gormx"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"github.com/micro-services-roadmap/kit-common/zapx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitV0(conf any, path ...string) {
	viperx.InitViperV2(conf, path...)
	zapx.InitZap()
	gormx.InitDB()
}

func Init(conf any, path ...string) (v *viper.Viper, l *zap.Logger, d *gorm.DB) {
	v = viperx.InitViperV2(conf, path...)
	l = zapx.InitZap()
	d = gormx.InitDB()

	return
}

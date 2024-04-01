package kg

import (
	"github.com/spf13/viper"
	"github.com/wordpress-plus/kit-common/viperx/vconfig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	C  vconfig.Server
	V  *viper.Viper
	L  *zap.Logger
	DB *gorm.DB
)

const (
	DbMysql = "mysql"
	DbPgsql = "pgsql"
)

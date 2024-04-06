package kg

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/wordpress-plus/kit-common/viperx/vconfig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	C     vconfig.Server
	V     *viper.Viper
	L     *zap.Logger
	DB    *gorm.DB
	REDIS *redis.Client
)

const (
	DbMysql = "mysql"
	DbPgsql = "pgsql"
)

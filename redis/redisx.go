package gormx

import (
	kg "github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/redis/initialize"
)

func InitRedis() {
	kg.REDIS = initialize.Redis()
}

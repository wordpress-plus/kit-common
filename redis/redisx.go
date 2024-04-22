package gormx

import (
	kg "github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/redis/initialize"
)

func InitRedis() {
	kg.REDIS = initialize.Redis()
}

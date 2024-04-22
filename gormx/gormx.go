package gormx

import (
	"github.com/micro-services-roadmap/kit-common/gormx/initialize"
	kg "github.com/micro-services-roadmap/kit-common/kg"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbType := kg.C.System.DbType
	switch dbType {
	case kg.DbMysql:
		kg.DB = initialize.GormMysql(true)
	case kg.DbPgsql:
		kg.DB = initialize.GormPgSQL(true)
	default:
		panic("unknown db type")
	}

	return kg.DB
}

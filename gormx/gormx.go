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
		kg.DB = initialize.GormMysql(kg.C.Mysql.Migration)
	case kg.DbPgsql:
		kg.DB = initialize.GormPgSQL(kg.C.Pgsql.Migration)
	default:
		panic("unknown db type")
	}

	return kg.DB
}

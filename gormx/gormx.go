package gormx

import (
	"github.com/wordpress-plus/kit-common/gormx/initialize"
	kg "github.com/wordpress-plus/kit-common/kg"
)

func InitDB() {
	dbType := kg.C.System.DbType
	switch dbType {
	case kg.DbMysql:
		kg.DB = initialize.GormMysql(true)
	case kg.DbPgsql:
		kg.DB = initialize.GormPgSQL(true)
	default:
		panic("unknown db type")
	}
}

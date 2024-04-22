package internal

import (
	"fmt"

	kg "github.com/micro-services-roadmap/kit-common/kg"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch kg.C.System.DbType {
	case "mysql":
		logZap = kg.C.Mysql.LogZap
	case "pgsql":
		logZap = kg.C.Pgsql.LogZap
	}
	if logZap {
		kg.L.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

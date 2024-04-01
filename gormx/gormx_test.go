package gormx

import (
	"fmt"
	kg "github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/viperx"
	"testing"
)

func init() {
	viperx.InitViper("config.yaml")
}

func TestInitDB(t *testing.T) {
	// 测试InitDB函数
	InitDB(kg.DbMysql)
	fmt.Printf("mysql: %v", kg.DB)
}

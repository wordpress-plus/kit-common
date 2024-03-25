package viperx

import (
	"fmt"
	"github.com/wordpress-plus/kit-logger/zapx/config"
	"testing"
)

func TestViper(_ *testing.T) {
	conf := &Conf{}
	v := Viper(conf, "config.yaml")
	fmt.Println(v)
}

type Conf struct {
	Zap    config.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System     `mapstructure:"system" json:"system" yaml:"system"`
}

type System struct {
	Env       string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	JasyptPwd string `mapstructure:"jasypt-pwd" json:"jasypt-pwd" yaml:"jasypt-pwd"` // 配置加解密值
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值

	// DbType 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DbType   string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	UseRedis bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
}

package vconfig

import (
	"github.com/wordpress-plus/kit-common/gormx/gconfig"
	"github.com/wordpress-plus/kit-common/zapx/config"
)

type Server struct {
	gconfig.DbServer `yaml:",inline" mapstructure:",squash"`

	Zap    config.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System     `mapstructure:"system" json:"system" yaml:"system"`
}

type System struct {
	gconfig.DbSys `yaml:",inline" mapstructure:",squash"`

	Env       string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	JasyptPwd string `mapstructure:"jasypt-pwd" json:"jasypt-pwd" yaml:"jasypt-pwd"` // 配置加解密值
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值
	UseRedis  bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`    // 使用redis
}

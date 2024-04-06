package kit

import (
	"github.com/wordpress-plus/kit-common/gormx"
	"github.com/wordpress-plus/kit-common/viperx"
	"github.com/wordpress-plus/kit-common/zapx"
)

func Init(conf any, path ...string) {
	viperx.InitViperV2(conf, path...)
	zapx.InitZap()
	gormx.InitDB()
}

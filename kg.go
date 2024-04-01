package kit

import (
	"github.com/wordpress-plus/kit-common/gormx"
	"github.com/wordpress-plus/kit-common/viperx"
	"github.com/wordpress-plus/kit-common/zapx"
)

func Init(path ...string) {
	viperx.InitViper(path...)
	zapx.InitZap()
	gormx.InitDB()
}

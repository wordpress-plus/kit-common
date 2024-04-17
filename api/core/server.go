package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wordpress-plus/kit-common/kg"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer(Router *gin.Engine) {

	address := fmt.Sprintf(":%d", kg.C.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	kg.L.Info("server run success on ", zap.String("address", address))
	kg.L.Error(s.ListenAndServe().Error())
}

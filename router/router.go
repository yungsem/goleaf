package router

import (
	"github.com/gin-gonic/gin"
	handler2 "github.com/yungsem/goleaf/handler"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// 分配 ID
	r.GET("/allocate", handler2.Allocate)
	// 监控数据
	r.GET("/info", handler2.Info)

	// 开启 pprof 监控
	//ginpprof.Wrap(r)
	return r
}

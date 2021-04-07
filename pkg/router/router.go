package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yungsem/goleaf/pkg/handler"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// 分配 ID
	r.GET("/allocate", handler.Allocate)
	// 监控数据
	r.GET("/info", handler.Info)

	return r
}

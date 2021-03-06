package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yungsem/goleaf/core/allocator"
	inits2 "github.com/yungsem/goleaf/inits"
	"github.com/yungsem/gotool/result"
	"net/http"
)

// Allocate 获取 ID
func Allocate(c *gin.Context) {

	bizTag := c.Query("bizTag")
	inits2.Log.Debug("收到请求，bizTag=%s", bizTag)

	nextId, err := allocator.AllocateId(bizTag)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success(nextId))
}

// Info 获取 buffer 信息
func Info(c *gin.Context) {
	bizTag := c.Query("bizTag")

	infoList := allocator.BizBufferMapInfo(bizTag)

	c.JSON(http.StatusOK, result.Success(infoList))
}

package api

import (
	"Dorm/serializer"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.BaseResponse{
		Status: 0000,
		Msg:    "pong",
	})
}

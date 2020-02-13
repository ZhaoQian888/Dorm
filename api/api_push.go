package api

import (
	"Dorm/cache"
	"Dorm/serializer"

	"github.com/gin-gonic/gin"
)

// AdminPush 管理redis缓冲。应交给服务器自动执行，但还为想好策略
func AdminPush(c *gin.Context) {
	err := cache.Push()
	if err != nil {
		c.JSON(200, serializer.BaseResponse{
			Status: 0022,
			Msg:    "cache 出错",
			Error:  err.Error(),
		})
	} else {
		c.JSON(200, serializer.BaseResponse{
			Status: 0,
			Msg:    "缓存成功",
		})
	}

}

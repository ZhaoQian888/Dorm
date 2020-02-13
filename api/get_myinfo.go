package api

import (
	"Dorm/serializer"

	"github.com/gin-gonic/gin"
)

// ReturnID 返回openid
func ReturnID(c *gin.Context) {
	c.JSON(200, serializer.BaseResponse{
		Status: 0,
		Data:   "ok",
	})
}

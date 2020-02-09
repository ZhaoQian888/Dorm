package api

import (
	"Dorm/serializer"
	"Dorm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RemainDorm GET请求，返回剩余宿舍信息
func RemainDorm(c *gin.Context) {
	var rm = service.RemainDormService{}
	var err error
	rm.Depart = c.Query("depart")
	rm.Gender, err = strconv.Atoi(c.Query("gender"))
	if err != nil {
		c.JSON(200, serializer.BaseResponse{
			Msg:    "参数序列化失败，gender错误",
			Status: 0001,
			Error:  err.Error(),
		})
	}
	rm.Size, err = strconv.Atoi(c.Query("size"))
	if err != nil {
		c.JSON(200, serializer.BaseResponse{
			Msg:    "参数序列化失败,size错误",
			Status: 0001,
			Error:  err.Error(),
		})
	}
	res := rm.Query()
	c.JSON(200, res)
}

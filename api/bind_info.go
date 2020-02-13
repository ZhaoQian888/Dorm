package api

import (
	"Dorm/serializer"
	"Dorm/service"

	"github.com/gin-gonic/gin"
)

// BindInfo  完成信息绑定
func BindInfo(c *gin.Context) {
	var b = service.BindInfo{}
	if err := c.ShouldBind(&b); err != nil {
		c.JSON(200, serializer.BaseResponse{
			Status: 231,
			Msg:    "序列化出错",
		})
	} else {
		opid, _ := c.Get("opid")
		id, _ := opid.(string)
		res := b.Binding(id)
		c.JSON(200, res)
	}

}

package api

import (
	"Dorm/serializer"
	"Dorm/service"

	"github.com/gin-gonic/gin"
)

// DistriDorm 分配宿舍
func DistriDorm(c *gin.Context) {
	var ds = service.DistrDormService{}
	err := c.ShouldBind(&ds)
	if err != nil {
		c.JSON(200, serializer.BaseResponse{
			Status: 0002,
			Msg:    "序列化失败",
			Error:  err.Error(),
		})
	} else {
		res, errs := ds.DistrDorm()
		if errs != nil {
			c.JSON(200, serializer.BaseResponse{
				Status: 0004,
				Data:   res,
				Error:  errs.Error(),
			})
		} else {
			c.JSON(200, serializer.BaseResponse{
				Status: 0000,
				Data:   res,
			})

		}

	}

}

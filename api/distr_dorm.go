package api

import (
	"Dorm/cache"
	"Dorm/serializer"
	"Dorm/service"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// DistriDorm 分配宿舍
func DistriDorm(c *gin.Context) {
	var ds = service.DistrDormService{}
	de, ok := c.Get("depart")
	if !ok {
		ds.Depart = ""
	} else {
		ds.Depart = de.(string)
	}
	ge, _ := c.Get("gender")

	ds.Gender = ge.(int)

	stu, _ := c.Get("stu_number")
	ds.StuNumber = stu.(int)
	size, ok := c.Get("size")
	if !ok {
		ds.Size = 0
	} else {
		ds.Size = size.(int)
	}
	res, errs := ds.DistrDorm()
	s := cache.REDISPOOL.Get()
	dgs, _ := redis.String(s.Do("get", "DGS"+res))
	if errs != nil {
		c.JSON(200, serializer.BaseResponse{
			Status: 0004,
			Data:   []string{res, dgs},
			Error:  errs.Error(),
		})
	} else {
		c.JSON(200, serializer.BaseResponse{
			Status: 0000,
			Data:   []string{res, dgs},
		})

	}

}

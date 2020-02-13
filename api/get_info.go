package api

import (
	"Dorm/cache"
	"Dorm/serializer"
	"Dorm/service"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// RemainBed GET请求，返回剩余宿舍信息
func RemainBed(c *gin.Context) {
	var rm = service.RemainBedService{}
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

// ToRemainBed 模糊查找
func ToRemainBed(c *gin.Context) {
	res := service.ToQuery()
	c.JSON(200, res)
}

// GetInfo 返回用户信息
func GetInfo(c *gin.Context) {
	a, _ := c.Get("opid")
	b, _ := a.(string)
	u, err := service.GetUserInfo(b)
	if err != nil {
		c.JSON(200, serializer.BaseResponse{
			Status: 0043,
			Error:  err.Error(),
			Data:   u,
		})
	} else {
		c.JSON(200, serializer.BaseResponse{
			Status: 0,
			Data:   u,
		})
	}

}

// WtChoose 是否选过宿舍
func WtChoose(c *gin.Context) {
	stu, ok := c.Get("stu_number")
	if !ok {
		c.JSON(200, serializer.BaseResponse{
			Status: 78,
			Msg:    "需要绑定",
		})
	} else {
		t, _ := stu.(int)
		s := cache.REDISPOOL.Get()
		q, err := redis.String(s.Do("get", t))
		if err != nil || q == "" {
			c.JSON(200, serializer.BaseResponse{
				Data:   q,
				Status: 33,
			})
		} else {
			dgs, _ := redis.String(s.Do("get", "DGS"+q))
			c.JSON(200, serializer.BaseResponse{
				Data:   []string{q, dgs},
				Status: 0,
			})
		}
	}

}

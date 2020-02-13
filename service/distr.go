package service

import (
	"Dorm/cache"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// DistrDormService 分配宿舍服务
type DistrDormService struct {
	Depart    string `form:"depart" json:"depart" binding:"max=4"`
	Gender    int    `form:"gender" json:"gender" binding:"max=1" `
	Size      int    `form:"size"   json:"size"   binding:"max=6"`
	StuNumber int    `form:"stu_number" json:"stu_number" binding:"required"`
}

// DistrDorm 完成分配宿舍
func (d *DistrDormService) DistrDorm() (string, error) {
	c := cache.REDISPOOL.Get()
	a, _ := redis.String(c.Do("get", d.StuNumber))
	if a != "" {
		return a, errors.New("不可重复选宿舍")
	}
	var info = cache.DormRequir{
		Depart: d.Depart,
		Gender: uint8(d.Gender),
		Size:   uint8(d.Size),
	}
	dorm, err := cache.Hit(info)
	if err != nil {
		return dorm, err
	}
	err = cache.Record(d.StuNumber, dorm)
	return dorm, err
}

// VaildCheck 验证信息可靠性
func (d *DistrDormService) VaildCheck() (bool, error) {
	return true, errors.New("ok")
}

package cache

import (
	"github.com/garyburd/redigo/redis"
)

// Record 在redis中将学生号和床位号对应记录
func Record(stu int, bed string) error {
	c := REDISPOOL.Get()
	_, err := c.Do("set", stu, bed)
	if err != nil {
		return err
	}
	_, err = redis.String(c.Do("set", bed, stu))
	return err
}

package cache

import (
	"github.com/garyburd/redigo/redis"
)

// Record 记录分配
func Record(stu int, bed string) error {
	c := REDISPOOL.Get()
	_, err := c.Do("set", stu, 1)
	if err != nil {
		return err
	}
	_, err = redis.String(c.Do("set", bed, stu))
	return err
}

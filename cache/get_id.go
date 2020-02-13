package cache

import (
	"Dorm/model"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// GetStuNumGen 缓存层查询
func GetStuNumGen(id string) (int, int, error) {
	c := REDISPOOL.Get()
	num, err := redis.String(c.Do("hget", id, "stunumber"))
	gen, err := redis.String(c.Do("hget", id, "gender"))
	if err != nil {
		return model.GetStuNumGen(id)
	}
	a, _ := strconv.Atoi(num)
	b, _ := strconv.Atoi(gen)
	return a, b, nil
}

package cache

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

// GetOpenID 查询redis中的 opid
func GetOpenID(uid interface{}) (string, error) {
	a, ok := uid.(string)
	if !ok {
		return "", errors.New("断言出错")
	}
	c := REDISPOOL.Get()
	opid, err := redis.String(c.Do("hget", a, "openid"))
	return opid, err
}

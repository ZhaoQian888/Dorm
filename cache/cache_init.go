package cache

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

// REDISPOOL 是一个redis连接池
var REDISPOOL *redis.Pool

// RedisPoolInit 用来初始化连接池
func redisPoolInit() {
	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_ADDR"))
		},
	}
	REDISPOOL = pool
	REDISPOOL.Get()
}

// Init 初始化所有cache操作
func Init() {
	redisPoolInit()
	cacheInit()
}

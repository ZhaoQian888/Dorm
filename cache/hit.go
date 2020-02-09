package cache

import (
	"math/rand"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// DormRequir 选宿舍要求
type DormRequir struct {
	Depart string
	Gender uint8
	Size   uint8
}

// Hit 用来命中合理的宿舍
func Hit(infos DormRequir) (string, error) {
	c := REDISPOOL.Get()
	var di int
	var gi int
	di = rand.Int() % 5
	gi = rand.Int() % 3
	if infos.Depart != "" {
		di = dmap[infos.Depart]
	}
	if infos.Size != 0 {
		gi = smap[infos.Size]
	}
	for i := di; i < di+5; i++ {
		for j := gi; j < gi+3; j++ {
			key := "D_" + depart[i%5] + "G_" + strconv.Itoa(int(infos.Gender)) + "S_" + strconv.Itoa(int(size[j%3]))
			re, err := redis.Int(c.Do("decr", key))
			if err != nil {
				return "", err
			}
			if re >= 0 {
				return redis.String(c.Do("get", re))
			}
		}
	}
	return "", nil
}

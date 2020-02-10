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
	var res string = ""
	di = rand.Int() % len(dmap)
	gi = rand.Int() % len(smap)
	if infos.Depart != "" {
		di = dmap[infos.Depart]
	}
	if infos.Size != 0 {
		gi = smap[infos.Size]
	}
	for i := di; i < di+len(dmap); i++ {
		for j := gi; j < gi+len(smap); j++ {
			key := "D_" + depart[i%len(dmap)] + "G_" + strconv.Itoa(int(infos.Gender)) + "S_" + strconv.Itoa(int(size[j%len(smap)]))
			re, err := redis.Int(c.Do("decr", key))
			if err != nil {

			}
			if re >= 0 {
				ress, err := redis.String(c.Do("get", key+strconv.Itoa(re)))
				res = ress
				if err == nil {
					break
				}
			}
		}
	}
	return res, nil
}

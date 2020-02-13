package cache

import (
	"Dorm/model"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

var depart = []string{"5B", "8F", "9D", "13E", "14A"}
var gender = []uint8{0, 1}
var size = []uint8{2, 3, 4, 5}
var dmap = map[string]int{
	depart[0]: 0,
	depart[1]: 1,
	depart[2]: 2,
	depart[3]: 3,
	depart[4]: 4,
}
var smap = map[uint8]int{
	size[0]: 0,
	size[1]: 1,
	size[2]: 2,
	size[3]: 3,
}

// CacheInit 将数据库中的数据缓存到redis中
func cacheInit() error {
	bedCache()
	departCache()
	return nil
}

// departCache
func departCache() error {
	c := REDISPOOL.Get()
	for i := 0; i < len(depart); i++ {
		a, err := strconv.Atoi(depart[i][:(len(depart[i]) - 1)])
		if err != nil {
			return err
		}
		num, err := model.GetDepartBedNumber(a, 1)
		if err != nil {
			return err
		}
		_, err = c.Do("set", depart[i]+"M", num)
		if err != nil {
			return err
		}
		num, err = model.GetDepartBedNumber(a, 0)
		if err != nil {
			return err
		}
		_, err = c.Do("set", depart[i]+"W", num)
		if err != nil {
			return err
		}
	}
	return nil
}

func bedCache() {
	c := REDISPOOL.Get()
	for i := 0; i < len(depart); i++ {
		for j := 0; j < len(gender); j++ {
			for n := 0; n < len(size); n++ {
				beds, _ := model.FindAllBed(depart[i], gender[j], size[n])
				c.Do("set", "D_"+depart[i]+"G_"+strconv.Itoa(int(gender[j]))+"S_"+strconv.Itoa(int(size[n])), len(beds))
				for m := 0; m < len(beds); m++ {
					c.Do("set", "D_"+depart[i]+"G_"+strconv.Itoa(int(gender[j]))+"S_"+strconv.Itoa(int(size[n]))+strconv.Itoa(m), beds[m].BedNumber)
					c.Do("set", beds[m].BedNumber, 1)
					c.Do("sadd", "beds", beds[m].BedNumber)
					c.Do("set", "DGS"+beds[m].BedNumber, "D_"+depart[i]+"G_"+strconv.Itoa(int(gender[j]))+"S_"+strconv.Itoa(int(size[n]))+strconv.Itoa(m))
				}
			}
		}
	}
}

// Push 将redis中数据缓冲至mysql
func Push() error {
	c := REDISPOOL.Get()
	beds, _ := redis.StringMap(c.Do("SMEMBERS", "beds"))
	var info []model.PushInfo
	var student uint64
	for i := 0; i < len(beds); i++ {
		student, _ = redis.Uint64(c.Do("get", beds[strconv.Itoa(i)]))
		info = append(info, model.PushInfo{
			BedNumber: beds[strconv.Itoa(i)],
			StuNumber: student,
		})
	}
	err := model.PushMysql(info)
	return err
}

package cache

import (
	"Dorm/model"
	"strconv"
)

// Init 初始化所有cache操作
func Init() {
	redisPoolInit()
	cacheInit()
}

// CacheInit 将数据库中的数据缓存到redis中
func cacheInit() error {
	departCache()
	unitCache()
	dormCache()
	return nil
}

// departCache 用来将所有的宿舍楼的床位个数映射到redis中
func departCache() {
	c := REDISPOOL.Get()
	depart := []int{5, 8, 9, 13, 14}
	for i := 0; i < 5; i++ {
		departbednumber, err := model.GetDepartBedNumber(depart[i])
		if err != nil {
			panic(err)
		}
		c.Do("set", "depart"+strconv.Itoa(depart[i]), departbednumber)
	}
}

// unitCache 用来将当前所有的单元楼的床位个数映射到redis中
func unitCache() {
	c := REDISPOOL.Get()
	unit, _ := model.FindAllUnit()
	for i := 0; i < len(unit); i++ {
		c.Do("set", "unit"+unit[i].UID, unit[i].Num)
	}
}

// dormCache
func dormCache() {
	c := REDISPOOL.Get()
	dorm, _ := model.FindAllDorm()
	for i := 0; i < len(dorm); i++ {
		c.Do("set", "dorm"+dorm[i].DormNumber, dorm[i].Num)
	}
}

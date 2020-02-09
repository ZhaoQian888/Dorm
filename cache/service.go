package cache

import (
	"strconv"

	"github.com/garyburd/redigo/redis"
)

//RemainService 为service 层提供查询服务
func RemainService(departs string, genders int, sizes int) (int, error) {
	c := REDISPOOL.Get()
	var res int
	switch {
	case departs == "" && genders == 2 && sizes == 0:
		for i := 0; i < len(depart); i++ {
			for j := 0; j < len(size); j++ {
				for n := 0; n < len(gender); n++ {
					key := "D_" + depart[i] + "G_" + strconv.Itoa(int(gender[n])) + "S_" + strconv.Itoa(int(size[j]))
					temp, errs := redis.Int(c.Do("get", key))
					if errs != nil {
						return 0, errs
					}
					res += temp
				}
			}
		}
		return res, nil
	case departs == "" && genders == 2 && sizes != 0:
		for i := 0; i < len(depart); i++ {
			for j := 0; j < len(gender); j++ {
				key := "D_" + depart[i] + "G_" + strconv.Itoa(int(gender[j])) + "S_" + strconv.Itoa(sizes)
				temp, errs := redis.Int(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				res += temp
			}
		}
		return res, nil
	case departs == "" && genders != 2 && sizes == 0:
		for i := 0; i < len(gender); i++ {
			for j := 0; j < len(size); j++ {
				key := "D_" + depart[i] + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(int(size[j]))
				temp, errs := redis.Int(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				res += temp
			}
		}
		return res, nil
	case departs != "" && genders == 2 && sizes == 0:
		for i := 0; i < len(gender); i++ {
			for j := 0; j < len(size); j++ {
				key := "D_" + departs + "G_" + strconv.Itoa(int(gender[i])) + "S_" + strconv.Itoa(int(size[j]))
				temp, errs := redis.Int(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				res += temp
			}
		}
		return res, nil
	case departs != "" && genders != 2 && sizes == 0:
		for i := 0; i < len(size); i++ {
			key := "D_" + departs + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(int(size[i]))
			temp, errs := redis.Int(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			res += temp
		}
		return res, nil
	case departs != "" && genders == 2 && sizes != 0:
		for i := 0; i < len(gender); i++ {
			key := "D_" + departs + "G_" + strconv.Itoa(int(gender[i])) + "S_" + strconv.Itoa(sizes)
			temp, errs := redis.Int(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			res += temp
		}
		return res, nil
	case departs == "" && genders != 2 && sizes != 0:
		for i := 0; i < len(depart); i++ {
			key := "D_" + depart[i] + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(sizes)
			temp, errs := redis.Int(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			res += temp
		}
		return res, nil
	default:
		key := "D_" + departs + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(sizes)
		temp, errs := redis.Int(c.Do("get", key))
		if errs != nil {
			return 0, errs
		}
		res += temp
		return res, nil
	}
}

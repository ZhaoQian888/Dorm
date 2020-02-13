package cache

import (
	"Dorm/model"
	"Dorm/serializer"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// DepartInfo 宿舍楼剩余情况
type DepartInfo struct {
	Depart string
	Gender int
	Size   int
	Num    int
}

// TotalRemainDorm 所有宿舍剩余情况
type TotalRemainDorm struct {
	AllDepart  []string
	DepartNum  int
	DepartInfo []DepartInfo
}

// TotalBed 返回宿舍剩余情况 为service 层提供查询服务
func TotalBed() serializer.BaseResponse {
	var res = TotalRemainDorm{}
	for i := 0; i < len(depart); i++ {
		res.AllDepart = append(res.AllDepart, depart[i])
		res.DepartNum++
		for j := 0; j < len(gender); j++ {
			for n := 0; n < len(size); n++ {
				nu, err := RemainBedService(depart[i], int(gender[j]), int(size[n]))
				if err != nil {
					return serializer.BaseResponse{
						Status: 0003,
						Msg:    "查询出错",
						Error:  err.Error(),
					}
				}
				res.DepartInfo = append(res.DepartInfo, DepartInfo{
					Depart: depart[i],
					Gender: int(gender[j]),
					Size:   int(size[n]),
					Num:    nu,
				})
			}
		}
	}
	return serializer.BaseResponse{
		Status: 0000,
		Msg:    "ok",
		Data:   res,
		Error:  "",
	}
}

//RemainBedService 为service 层提供查询服务（精确查询）
func RemainBedService(departs string, genders int, sizes int) (int, error) {
	c := REDISPOOL.Get()
	var res int = 0
	switch {
	case departs == "" && genders == 2 && sizes == 0:
		for i := 0; i < len(depart); i++ {
			for j := 0; j < len(size); j++ {
				for n := 0; n < len(gender); n++ {
					key := "D_" + depart[i] + "G_" + strconv.Itoa(int(gender[n])) + "S_" + strconv.Itoa(int(size[j]))
					temp, errs := redis.String(c.Do("get", key))
					if errs != nil {
						return 0, errs
					}
					a, _ := strconv.Atoi(temp)
					res += a
				}
			}
		}
		return res, nil
	case departs == "" && genders == 2 && sizes != 0:
		for i := 0; i < len(depart); i++ {
			for j := 0; j < len(gender); j++ {
				key := "D_" + depart[i] + "G_" + strconv.Itoa(int(gender[j])) + "S_" + strconv.Itoa(sizes)
				temp, errs := redis.String(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				a, _ := strconv.Atoi(temp)
				res += a
			}
		}
		return res, nil
	case departs == "" && genders != 2 && sizes == 0:
		for i := 0; i < len(gender); i++ {
			for j := 0; j < len(size); j++ {
				key := "D_" + depart[i] + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(int(size[j]))
				temp, errs := redis.String(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				a, _ := strconv.Atoi(temp)
				res += a
			}
		}
		return res, nil
	case departs != "" && genders == 2 && sizes == 0:
		for i := 0; i < len(gender); i++ {
			for j := 0; j < len(size); j++ {
				key := "D_" + departs + "G_" + strconv.Itoa(int(gender[i])) + "S_" + strconv.Itoa(int(size[j]))
				temp, errs := redis.String(c.Do("get", key))
				if errs != nil {
					return 0, errs
				}
				a, _ := strconv.Atoi(temp)
				res += a
			}
		}
		return res, nil
	case departs != "" && genders != 2 && sizes == 0:
		for i := 0; i < len(size); i++ {
			key := "D_" + departs + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(int(size[i]))
			temp, errs := redis.String(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			a, _ := strconv.Atoi(temp)
			res += a
		}
		return res, nil
	case departs != "" && genders == 2 && sizes != 0:
		for i := 0; i < len(gender); i++ {
			key := "D_" + departs + "G_" + strconv.Itoa(int(gender[i])) + "S_" + strconv.Itoa(sizes)
			temp, errs := redis.String(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			a, _ := strconv.Atoi(temp)
			res += a
		}
		return res, nil
	case departs == "" && genders != 2 && sizes != 0:
		for i := 0; i < len(depart); i++ {
			key := "D_" + depart[i] + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(sizes)
			temp, errs := redis.String(c.Do("get", key))
			if errs != nil {
				return 0, errs
			}
			a, _ := strconv.Atoi(temp)
			res += a
		}
		return res, nil
	default:
		key := "D_" + departs + "G_" + strconv.Itoa(genders) + "S_" + strconv.Itoa(sizes)
		temp, errs := redis.String(c.Do("get", key))
		if errs != nil {
			return 0, errs
		}
		a, _ := strconv.Atoi(temp)
		res += a
		return res, nil
	}
}

// DormInfo 提供宿舍信息
type DormInfo struct {
	DromNumber string
	Size       int
	UnUseSize  int
}

// GetUser 先查redis 命中返回。 未命中去查mysql  b是opid
func GetUser(b string) (model.Student, error) {
	c := REDISPOOL.Get()
	stunumber, err := redis.String(c.Do("hget", b, "stunumber"))
	if err != nil {
		user, err := model.GetStudent(b)
		if err != nil {
			return user, err
		}
		err = CachStudent(user.StuName, user.StuNmber, user.StuTele, user.Gender, user.StuWeChat, user.Region, user.Class)
		return user, err
	}
	num, err := strconv.Atoi(stunumber)
	if err != nil {
		return model.Student{}, err
	}
	stuname, err := redis.String(c.Do("hget", b, "stuname"))
	if err != nil {
		return model.Student{}, err
	}
	stutel, err := redis.String(c.Do("hget", b, "stutel"))
	if err != nil {
		return model.Student{}, err
	}
	tel, err := strconv.Atoi(stutel)
	if err != nil {
		return model.Student{}, err
	}
	gender, err := redis.String(c.Do("hget", b, "gender"))
	if err != nil {
		return model.Student{}, err
	}
	ge, err := strconv.Atoi(gender)
	cla, err := redis.String(c.Do("hget", b, "class"))
	if err != nil {
		return model.Student{}, err
	}
	reg, err := redis.String(c.Do("hget", b, "region"))
	return model.Student{
		Region:    reg,
		Class:     cla,
		StuNmber:  uint64(num),
		StuName:   stuname,
		StuTele:   uint64(tel),
		Gender:    uint8(ge),
		StuWeChat: b,
	}, nil
}

// CachStudent 将数据库中的模型映射到redis中
func CachStudent(stuname string, stunumber uint64, stutel uint64, gender uint8, wechat string, region string, class string) error {

	c := REDISPOOL.Get()
	_, err := c.Do("hset", wechat, "stuname", stuname)
	if err != nil {
		return err
	}
	_, err = c.Do("hset", wechat, "stunumber", stunumber)
	if err != nil {
		return err
	}
	_, err = c.Do("hset", wechat, "gender", gender)
	if err != nil {
		return err
	}
	_, err = c.Do("hset", wechat, "stutel", stutel)
	if err != nil {
		return err
	}
	_, err = c.Do("hset", wechat, "region", region)
	if err != nil {
		return err
	}
	_, err = c.Do("hset", wechat, "class", class)
	c.Close()
	return nil

}

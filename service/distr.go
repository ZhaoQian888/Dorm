package service

import "Dorm/cache"

// DistrDorm 完成分配宿舍
func DistrDorm(depart string, gender int, size int) (string, error) {
	var info = cache.DormRequir{
		Depart: depart,
		Gender: uint8(gender),
		Size:   uint8(size),
	}
	dorm, err := cache.Hit(info)
	return dorm, err

}

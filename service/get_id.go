package service

import "Dorm/cache"

// GetStuNumGen 通过opid找到学生的stunumber
func GetStuNumGen(id string) (int, int, error) {
	num, gen, err := cache.GetStuNumGen(id)
	return num, gen, err
}

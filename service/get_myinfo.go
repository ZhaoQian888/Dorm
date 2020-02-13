package service

import "Dorm/cache"

// GetUserInfo 返回用户信息
func GetUserInfo(b string) (interface{}, error) {
	return cache.GetUser(b)
}

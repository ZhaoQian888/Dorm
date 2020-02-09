package service

import "Dorm/cache"

// Reorganize 整理数据，将所有data push到mysql中
func Reorganize() error {
	err := cache.Push()
	return err
}

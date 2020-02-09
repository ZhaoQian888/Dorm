package service

import (
	"Dorm/cache"
	"Dorm/serializer"
)

// RemainDormService 剩余床位查询
type RemainDormService struct {
	Depart string `form:"depart" json:"depart" binding:""`
	Gender int    `form:"gender" json:"gender" binding:""`
	Size   int    `form:"size"   json:"size"   binding:""`
}

// Query 返回剩余宿舍信息
func (r *RemainDormService) Query() serializer.BaseResponse {
	num, _ := cache.RemainService(r.Depart, r.Gender, r.Size)
	return serializer.BaseResponse{
		Status: 0000,
		// Error:  err.Error(),
		Data: num,
	}
}

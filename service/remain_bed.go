package service

import (
	"Dorm/cache"
	"Dorm/serializer"
)

// RemainBedService 剩余床位查询
type RemainBedService struct {
	Depart string `form:"depart" json:"depart" binding:""`
	Gender int    `form:"gender" json:"gender" binding:""`
	Size   int    `form:"size"   json:"size"   binding:""`
}

// Query 返回剩余宿舍信息
func (r *RemainBedService) Query() serializer.BaseResponse {

	num, err := cache.RemainBedService(r.Depart, r.Gender, r.Size)
	if err != nil {
		return serializer.BaseResponse{
			Status: 000,
			Error:  err.Error(),
			Data:   num,
		}
	}
	return serializer.BaseResponse{
		Status: 0000,
		Data:   num,
	}

}

// ToQuery 普通查找
func ToQuery() serializer.BaseResponse {
	return cache.TotalBed()
}

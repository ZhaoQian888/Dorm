package service

import (
	"Dorm/model"
	"Dorm/serializer"
	"fmt"
)

// BindInfo 绑定信息
type BindInfo struct {
	StuName   string   `json:"name"      form:"name"`
	StuNumber int      `json:"stunumber" form:"stunumber"`
	Gender    int      `json:"gender"    form:"gender"`
	Region    []string `json:"region"    form:"region"`
	Tel       int      `json:"phone"     form:"phone"`
	Cla       string   `json:"classes"   form:"classes"`
}

// Binding  绑定数据库
func (b *BindInfo) Binding(id string) serializer.BaseResponse {
	var r = ""
	for i := 0; i < len(b.Region); i++ {
		r = r + b.Region[i]
	}
	fmt.Print(b)
	err := model.SetStudent(b.StuNumber, b.StuName, id, b.Tel, b.Gender, r, b.Cla)
	if err != nil {
		return serializer.BaseResponse{
			Status: 445,
			Msg:    "绑定出错",
			Error:  err.Error(),
		}
	}
	return serializer.BaseResponse{
		Status: 0,
		Msg:    "ok",
	}
}

package model

import "github.com/jinzhu/gorm"

// StudentBed 选宿舍结果
type StudentBed struct {
	//完整性约束 未在数据库实现，请程序猿人工实现- -。
	gorm.Model
	Student  Student `gorm:"foreign_key:StuRefer"`
	StuRefer uint64  `gorm:"unique;not null"`
	Bed      Bed     `gorm:"foreign_key:BedRefer"`
	BedRefer string  `gorm:"unique;not null"`
}

func pushSB(info PushInfo) error {
	var sb = StudentBed{
		StuRefer: info.StuNumber,
		BedRefer: info.BedNumber,
	}
	err := MYSQL.Create(&sb).Error
	return err
}

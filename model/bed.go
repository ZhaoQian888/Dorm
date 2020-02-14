package model

// Bed 床位
type Bed struct {
	// 床位号	宿舍号+床位 13E34133
	BedNumber string `gorm:"primary_key type:varchar(100)"`
	// 是否已使用 1代表未使用。0代表使用
	Remain uint8
	// 宿舍  如：13E3414
	Dorm      Dorm   `gorm:"foreignkey:DormRefer"`
	DormRefer string `gorm:"not null"`
}

// BedNum 传递床位信息
type BedNum struct {
	BedNumber string
	Remain    uint8
}

// FindAllBed 返回床位信息数组
func FindAllBed(depart string, gender uint8, size uint8) ([]BedNum, error) {
	var beds []Bed
	var bednums []BedNum
	err := MYSQL.Joins("JOIN dorms on dorms.dorm_number= beds.dorm_refer").
		Joins("JOIN units on units.uid=dorms.unit_refer").
		Joins("JOIN departs on departs.id=units.depart_refer").
		Where("dorms.gender=? and dorms.size=? and departs.id=? and remain=1", gender, size, depart).
		Find(&beds).
		Error
	if err != nil {
		return bednums, err
	}
	for i := 0; i < len(beds); i++ {
		bednums = append(bednums, BedNum{
			BedNumber: beds[i].BedNumber,
			Remain:    beds[i].Remain,
		})
	}
	return bednums, nil

}

func pushBed(bedNumber PushInfo) error {
	var bed Bed
	bed.BedNumber = bedNumber.BedNumber
	err := MYSQL.Model(&bed).Update("remain", "1").Error
	return err
}

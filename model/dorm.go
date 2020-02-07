package model

// Dorm 宿舍
type Dorm struct {
	// 宿舍号  13号宿舍楼E3414宿舍：13E3414
	DormNumber string `gorm:"primary_key"`
	// 单元
	Unit      Unit `gorm:"foreignkey:UnitRefer"`
	UnitRefer string
	// 楼层
	Floor uint8
	// 宿舍床位的个数
	Size uint8
	// 空的床位个数
	UnUseSize uint8
	//性别 0代表女生宿舍，1代表男生宿舍
	Gender uint8
	// 其他信息
	Info string
}

// DormNum 传递宿舍剩余床位信息
type DormNum struct {
	DormNumber string
	Num        uint8
}

// FindAllDorm 返回宿舍剩余床位信息
func FindAllDorm() ([]DormNum, error) {
	var dorms []Dorm
	var dormnums []DormNum
	var beds []Bed
	var c uint8
	err := MYSQL.Find(&dorms).Error
	if err != nil {
		return dormnums, err
	}
	for i := 0; i < len(dorms); i++ {
		err = MYSQL.Joins("JOIN dorms on dorms.dorm_number=beds.dorm_refer").
			Where("dorms.dorm_number=?", dorms[i]).
			Find(&beds).
			Count(&beds).
			Error
		dormnums = append(dormnums, DormNum{
			DormNumber: dorms[i].DormNumber,
			Num:        c,
		})
	}
	return dormnums, err
}

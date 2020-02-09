package model

// Unit 是一个单元
type Unit struct {
	// 单元号
	UID string `gorm:"primary_key"`
	// 宿舍楼
	Depart      Depart `gorm:"foreignkey:DepartRefer"`
	DepartRefer uint8
}

// UnitNum 用来传递Uint剩余床位信息
type UnitNum struct {
	UID string
	Num uint64
}

// FindAllUnit 返回所有的单元的信息的一个切片。
func FindAllUnit() ([]UnitNum, error) {
	var units []Unit
	var res []UnitNum
	var beds []Bed
	var c uint64
	err := MYSQL.Find(&units).Error
	if err != nil {
		return res, err
	}
	for i := 0; i < len(units); i++ {
		err = MYSQL.Joins("JOIN dorms on beds.dorm_refer=dorms.dorm_number").
			Joins("JOIN units on dorms.unit_refer=units.uid").
			Where("units.uid=? and beds.remain=?", units[i].UID, 1).
			Find(&beds).
			Count(&c).Error
		res = append(res, UnitNum{
			UID: units[i].UID,
			Num: c,
		})
	}
	return res, err
}

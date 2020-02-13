package model

// Depart 宿舍楼
type Depart struct {
	//宿舍楼号
	ID   uint8 `gorm:"primary_key;auto_increment:false"`
	Info string
}

// GetDepartBedNumber 获取宿舍楼的剩余床位数目
func GetDepartBedNumber(num int, gender int) (int, error) {
	var count int
	var beds []Bed
	err := MYSQL.Joins("JOIN dorms on beds.dorm_refer=dorms.dorm_number").
		Joins("JOIN units on units.uid=dorms.unit_refer").
		Joins("JOIN departs on departs.id=units.depart_refer").
		Where("depart_refer=? and gender=? and remain=0", num, gender).
		Find(&beds).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

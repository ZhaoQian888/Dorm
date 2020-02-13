package model

import "errors"

// Student 学生
type Student struct {
	//学号
	StuNmber uint64 `gorm:"primary_key;auto_increment:false"`
	//学生姓名
	StuName string
	//学生微信号
	StuWeChat string
	//学生手机号
	StuTele uint64
	// 性别 0女 1 男
	Gender uint8
	// 地区
	Region string
	// 学苑
	Class string
}

// GetStudent 查询数据库
func GetStudent(b string) (Student, error) {
	var s = Student{}
	var c = 0
	err := MYSQL.Where("stu_we_chat=?", b).First(&s).Count(&c).Error
	if err != nil {
		return s, err
	}
	if c == 0 {
		return s, errors.New("查无此人")
	}
	return s, nil
}

// SetStudent mysql 录入数据
func SetStudent(stunum int, stuname string, wechat string, tel int, gen int, reg string, cla string) error {
	var s = Student{
		StuNmber:  uint64(stunum),
		StuName:   stuname,
		StuWeChat: wechat,
		StuTele:   uint64(tel),
		Gender:    uint8(gen),
		Region:    reg,
		Class:     cla,
	}
	err := MYSQL.Create(&s).Error
	return err
}

// GetStuNumGen 查询数据库
func GetStuNumGen(id string) (int, int, error) {
	var s = Student{}
	var c = 0
	err := MYSQL.Where("stu_we_chat=?", id).First(&s).Count(&c).Error
	if err != nil {
		return 0, 0, err
	}
	if c == 0 {
		return 0, 0, errors.New("未绑定")
	}
	return int(s.StuNmber), int(s.Gender), nil
}

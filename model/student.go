package model

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
}

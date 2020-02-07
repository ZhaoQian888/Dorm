package model

// Bed 床位
type Bed struct {
	// 床位号	宿舍号+床位 13E34133
	BedNumber string `gorm:"primary_key"`
	// 是否已使用 1代表未使用。0代表使用
	Remain uint8
	// 宿舍  如：13E3414
	Dorm      Dorm   `gorm:"foreignkey:DormRefer"`
	DormRefer string `gorm:"not null"`
}

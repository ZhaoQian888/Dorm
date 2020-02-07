package model

func migration() {

	MYSQL.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Bed{}).
		AutoMigrate(&Dorm{}).
		AutoMigrate(&Unit{}).
		AutoMigrate(&Depart{}).
		AutoMigrate(&Student{}).
		AutoMigrate(&StudentBed{})
	tryBedModel()
}

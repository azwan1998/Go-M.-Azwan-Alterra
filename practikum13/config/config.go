package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// const DSN = "root:@/toko_buku?charset=utf8&parseTime=True&loc=Local"

func InitDB() {
	const DSN = "root:@tcp(192.168.32.240:3306)/toko_buku?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// func InitialMigration() {
// 	DB.AutoMigrate(&Buku{})
// }

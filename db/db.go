package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	dsn := "root:guaika1223@tcp(127.0.0.1:3306)/jack?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

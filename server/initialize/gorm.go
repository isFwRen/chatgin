package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IniMysql() *gorm.DB {
	var db *gorm.DB
	// https://github.com/go-sql-driver/mysql
	dsn := "root:root@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connection database")
	}
	return db
}

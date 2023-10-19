package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("===读取错误===")
	}
}
func InitMysql() *gorm.DB {
	//自定义日志末班，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,
		},
	)
	var db *gorm.DB
	// https://github.com/go-sql-driver/mysql
	//dsn := "root:root@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	// 迁移 schema
	//db.AutoMigrate(&model.UserBasic{})
	if err != nil {
		panic("failed to connection database")
	}
	return db
}

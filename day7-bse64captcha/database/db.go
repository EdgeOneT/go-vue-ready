package database

import (
	"fmt"
	"go-vue-ready/day7-bse64captcha/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.Cfg
	var err error
	username := cfg.Database.User
	password := cfg.Database.Password
	host := cfg.Database.Host
	port := cfg.Database.Port
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/admin-day2?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connection successful")
}

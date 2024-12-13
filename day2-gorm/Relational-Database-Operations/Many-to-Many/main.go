package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

var DB *gorm.DB

func initDB() {
	var err error
	dsn := "root:2021@tcp(127.0.0.1:3306)/admin-day2?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// 自动迁移
	DB.AutoMigrate(&User{}, &Role{})
	log.Println("Database connection successful")
}

func main() {
	initDB()
}

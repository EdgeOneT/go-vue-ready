package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User1 struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	ID     uint `gorm:"primaryKey"`
	Amount float64
	UserID uint
	User   User1 `gorm:"foreignKey:UserID"`
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
	DB.AutoMigrate(&User1{}, &Order{})
	log.Println("Database connection successful")
}

func main() {
	initDB()
}

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func initDB() {
	var err error
	dsn := "root:2021@tcp(127.0.0.1:3306)/admin-day2?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// 自动迁移
	//DB.AutoMigrate(&User{})
	log.Println("Database connection successful")
}

// 增
func Creat() {
	newUser := User{Name: "John", Email: "johndoe@example.com"}
	result := DB.Create(&newUser)
	if result.Error != nil {
		log.Fatal("Failed to insert user:", result.Error)
	}
	log.Println("New user created:", newUser)
}

// 查
func Read() {
	var user User
	result := DB.First(&user, 1) //根据ID查找
	if result.Error != nil {
		log.Fatal("Failed to find user:", result.Error)
	}
	log.Println("User found:", user)
}

// 改
func Update() {
	var user User
	result := DB.First(&user, 2)
	if result.Error != nil {
		log.Fatal("User not found:", result.Error)
	}
	user.Name = "Jane"
	DB.Save(&user)
	log.Println("User updated:", user)
}

// 删
func Delete() {
	var user User
	result := DB.First(&user, 2)
	if result.Error != nil {
		log.Fatal("User not found:", result.Error)
	}
	DB.Delete(&user)
	log.Println("User deleted:", user)
}
func main() {
	initDB()
	//Creat()

	//Read()

	//Update()

	Delete()
}

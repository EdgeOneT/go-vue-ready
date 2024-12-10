package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"'`
	Name      string `gorm:"size:100"`
	CreatedAt time.Time
}

type Order struct {
	ID     uint `gorm:"primaryKey"`
	Amount float64
	UserID uint
	user   User `gorm:"foreignKey:UserID"`
}

var DB *gorm.DB

func initDB() {
	var err error
	dsn := "root:2021@tcp(127.0.0.1:3306)/admin-day2?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	//DB.AutoMigrate(&User{}, &Order{})
}

func Creat() {
	user := User{Name: "小明"}
	order1 := Order{Amount: 121.5, user: user}
	order2 := Order{Amount: 999.9, user: user}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to insert user:", result.Error)
	}
	result1 := DB.Create(&order1)
	if result1.Error != nil {
		log.Fatal("Failed to insert user:", result.Error)
	}
	result2 := DB.Create(&order2)
	if result2.Error != nil {
		log.Fatal("Failed to insert user:", result.Error)
	}
	log.Println("添加成功")
}

func Delete(id int) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		log.Fatal("User not found:", result.Error)
	}
	DB.Delete(&user)
	log.Println("User deleted:", user)
}

func Update(id int, name string) {
	tx := DB.Begin()
	var user User
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		log.Fatal("Transaction failed:", err)
	}
	user.Name = name
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		log.Fatal("Transaction failed:", err)
	}
	tx.Commit()
	log.Println("User updated:", user)
}

func Read(id int) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		log.Fatal("User not found:", result.Error)
	}
	log.Println("User found:", user)
}

func main() {
	initDB()
	//Creat()
	Update(1, "小天")

}

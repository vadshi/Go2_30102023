package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID      uint
	Name    string
	Email   string
	Laptops []Laptop
}

type Laptop struct {
	ID           uint
	SerialNumber string
	UserID       uint
}

func main() {
	// SQLite does not support foreign key constraints
	db, err := gorm.Open(sqlite.Open("/tmp/example03.db"),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{}, &Laptop{})
	if err != nil {
		panic(err)
	}

	laptops := []Laptop{{SerialNumber: "sn0000001"}, {SerialNumber: "sn0000002"}}
	u := User{
		Name:    "User",
		Email:   "user@gmail.com",
		Laptops: laptops,
	}
	db.Create(&u)
	fmt.Println("Created", u)
	var recovered User
	db.First(&recovered)
	fmt.Println("Recovered without preload", recovered)
	recovered = User{}
	// Заполняем ссылки по внешнему ключу
	db.Preload("Laptops").First(&recovered)
	fmt.Println("Recovered with preload", recovered)
}

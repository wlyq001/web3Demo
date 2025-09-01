package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person1 struct {
	gorm.Model
	Name      string
	Age       int
	CompanyId uint
	Company   Company
}

type Person2 struct {
	gorm.Model
	Name       string
	Age        int
	CompanyKey uint
	Company    Company `gorm:"foreignKey:CompanyKey"`
}

type Person3 struct {
	gorm.Model
	Name       string
	Age        int
	CompanyKey uint
	Company    Company `gorm:"foreignKey:CompanyKey;references:CompanyCode"`
}

type Company struct {
	gorm.Model
	Name        string
	CompanyCode string `gorm:"unique"`
}

func main() {
	dsn := "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&Company{})
	//db.AutoMigrate(&Person1{})

	//db.AutoMigrate(&Company{})
	//db.AutoMigrate(&Person2{})

	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Person3{})
}

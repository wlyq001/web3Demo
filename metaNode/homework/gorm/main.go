package main

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    uint `gorm:"primaryKey"`
	Name  string
	Age   uint8
	Grade string
}

type Account struct {
	Id      uint `gorm:"primaryKey"`
	Balance uint
}

type Transaction struct {
	Id            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        uint
	Account1      Account `gorm:"foreignKey:FromAccountId"`
	Account2      Account `gorm:"foreignKey:ToAccountId"`
}

func main() {
	dsn := "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.Debug().AutoMigrate(&User{})

	//userZS := User{
	//	Name:  "李四",
	//	Age:   13,
	//	Grade: "一年级",
	//}
	//
	//db.Create(&userZS)

	//result := []User{}
	//db.Model(&User{}).Where("age > ?", 20).Find(&User{}).Scan(&result)
	//fmt.Println(result)

	//db.Model(&User{}).Where("name = ?", "张三").Update("grade", "四年级")

	//db.Model(&User{}).Where("age < ?", 15).Delete(&User{})

	//db.AutoMigrate(&Account{}, &Transaction{})

	db.Transaction(func(tx *gorm.DB) error {
		var accountA Account
		tx.Model(&Account{}).Where("id = 3").Find(&accountA)
		if accountA.Balance < 100 {
			return errors.New("balance is less than 100")
		}
		var accountB Account
		tx.Model(&Account{}).Where("id = 2").Find(&accountB)
		accountA.Balance -= 100
		accountB.Balance += 100
		updatesA := tx.Model(&Account{}).Updates(accountA)
		updatesB := tx.Model(&accountB).Updates(accountB)

		if updatesA.Error != nil || updatesB.Error != nil {
			return updatesA.Error
		}

		transaction := Transaction{
			FromAccountId: accountA.Id,
			ToAccountId:   accountB.Id,
			Amount:        100,
		}
		createT := tx.Create(&transaction)
		if createT.Error != nil {
			return createT.Error
		}
		return nil
	})
}

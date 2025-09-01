package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	Name         string
	Email        *string
	CreatedAt    time.Time
	BirthDate    *time.Time
	MemberNumber sql.NullString
}

type Author struct {
	User User   `gorm:"embedded;embeddedprefix:user_"`
	Book string `gorm:"column:novel_book"`
}

func main() {
	dsn := "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&User{})

	//user.MemberNumber.Valid = true
	users := []*User{&User{Name: "王李轶群", MemberNumber: sql.NullString{String: "1", Valid: true}}, &User{}}
	db.Create(users)

	db.AutoMigrate(Author{})
}

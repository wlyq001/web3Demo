package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Account struct {
	Id      uint `gorm:"primaryKey"`
	Balance uint
}

func (ac *Account) Scan(src any) error {
	var buf []byte
	switch v := src.(type) {
	case []byte:
		buf = v
	case string:
		buf = []byte(v)
	default:
		return fmt.Errorf("invalid type: %T", src)
	}

	err := json.Unmarshal(buf, ac)
	return err
}

func (ac Account) Value() (driver.Value, error) {
	v, err := json.Marshal(ac)
	return string(v), err
}

func initDB() (err error) {
	dsn := "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func queryRowDemo() {
	sqlStr := "select balance from accounts where balance > 0"
	var result []Account
	err := db.Get(&result, sqlStr)
	if err != nil {
		fmt.Printf("query row failed, err:%v\n", err)
		return
	}
	fmt.Println(result)

}

func main() {
	initDB()
	queryRowDemo()
}

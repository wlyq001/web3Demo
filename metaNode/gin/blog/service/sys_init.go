package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ENGINE *gin.Engine

func Init() {
	// 初始化数据库连接
	var dsn = "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 初始化gin
	ENGINE = gin.Default()
}

func RunGin() {
	err := ENGINE.Run("localhost:8080")
	if err != nil {
		return
	}
}

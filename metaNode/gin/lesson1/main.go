package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	engine.GET("/user", func(c *gin.Context) {
		user := User{
			Username: "admin",
			Password: "123456",
		}
		//c.JSON(200, &user)
		c.XML(200, &user)
	})

	grp1 := engine.Group("/v1")
	grp2 := engine.Group("/v2")

	grp1.Any("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "www.baidu.com")
		//c.JSON(200, "v1 hello world")
	})

	grp2.GET("/", func(c *gin.Context) {
		user := User{
			Username: "admin",
			Password: "123456",
		}
		c.JSON(200, &user)
	})

	err := engine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

package middleware

import (
	"time"

	"blog/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Path  string        // 访问路径
	Query string        // 携带query
	Body  string        // 携带body数据
	IP    string        // ip地址
	Error string        // 错误
	Cost  time.Duration // 花费时间
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := &Log{}
		defer func() {
			if err := recover(); err != nil {
				log.Error = err.(error).Error()
			}
		}()
		start := time.Now()
		log.Path = c.Request.URL.Path
		log.Query = c.Request.URL.RawQuery
		data, _ := c.GetRawData()
		log.Body = string(data)
		log.IP = c.ClientIP()
		c.Next()
		log.Cost = time.Since(start)
		service.DB.Create(log)
	}
}

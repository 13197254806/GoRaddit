//package helloworld

package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	})

	r.GET("/read", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("site_cookie")
		if err != nil {
			// 直接返回cookie值
			c.String(200, data)
			return
		}
		c.String(200, "not found!")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func main() {
	r := gin.Default()

	redirect(r)

	r.Run(":8111")
}

func redirect(r *gin.Engine) {
	// 支持内部外部重定向
	r.GET("/redirect", func(c *gin.Context) {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()

	// 服务端要给客户端cookie
	//r.GET("/cookie", func(c *gin.Context) {
	//	// 获取客户端是否携带cookie
	//	cookie, err := c.Cookie("sgm")
	//	if err != nil {
	//		cookie = "NotSet"
	//		// 设置cookie
	//		// maxAge int, 单位 s
	//		// path, cookie 所在目录
	//		// domain string, 域名
	//		// secure, 是否只能通过https访问
	//		// httpOnly bool 是否允许别人通过js获取cookie
	//		c.SetCookie("sgm", "value_cookie", 60, "/", "localhost", false, true)
	//	}
	//	fmt.Printf("cookie的值是：%s\n", cookie)
	//})

	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("sgm", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	r.Run(":8111")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Request.Cookie("sgm"); err == nil {
			if cookie.Value == "123" {
				c.Next()
				return
			}
		}

		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 如果验证不通过不在调用后续的函数处理
		c.Abort()
		return
	}
}

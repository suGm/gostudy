package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 定义计时中间件
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时:", since)
}

func main() {
	r := gin.Default()
	// 注册中间件
	//r.Use(MiddleWare())
	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	// {} 为了代码规范
	//{
	//	r.GET("middleware", func(c *gin.Context) {
	//		// 取值
	//		req, _ := c.Get("request")
	//		fmt.Println("request:", req)
	//		// 页面接收
	//		c.JSON(200, gin.H{"request": req})
	//	})
	//
	//	// 根路由后面是定义的局部中间件
	//	r.GET("middleware2", MiddleWare(), func(c *gin.Context) {
	//		// 取值
	//		req, _ := c.Get("request")
	//		fmt.Println("request:", req)
	//		// 页面接收
	//		c.JSON(200, gin.H{"request": req})
	//	})
	//}

	r.Run(":8111")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(time.Second * 3)
}

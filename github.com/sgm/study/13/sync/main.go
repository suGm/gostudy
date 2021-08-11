package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 同步异步
func main() {
	r := gin.Default()

	fooSync(r)

	r.Run(":8111")
}

// 1、goroutine机制可以方便地实现异步处理
// 2、另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
func fooSync(r *gin.Engine) {
	// 异步
	r.GET("/loginAsync", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 同步
	r.GET("/loginSync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步执行:" + c.Request.URL.Path)
	})
}

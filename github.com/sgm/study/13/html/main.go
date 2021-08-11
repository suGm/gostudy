package main

import "github.com/gin-gonic/gin"

// html渲染
func main() {
	// 1、创建路由
	// 默认使用了2个中间件Logger(),Recovery()
	r := gin.Default()

	setTitle(r)

	r.Run(":8111")
}

func setTitle(r *gin.Engine) {
	// 加载模板文件
	r.LoadHTMLGlob("./*")
	// r.LoadHTMLFiles("index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终json将title替换
		c.HTML(200, "index.tmpl", gin.H{"title": "sgm"})
	})
}

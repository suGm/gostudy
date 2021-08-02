package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// gin的helloWorld

type Login struct {
	// binding:"required" 修饰的字段，若接收为控制，则报错，是必选字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1、创建路由
	// 默认使用了2个中间件Logger(),Recovery()
	r := gin.Default()
	// 创建不带中间件的路由
	//r := gin.New()
	// 2、绑定路由规则
	// gin.Context, 封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/welcome", func(c *gin.Context) {
		// defaultQuery 第二个参数是默认值
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	r.POST("/form", func(c *gin.Context) {
		// 表单参数设置默认值
		type1 := c.DefaultPostForm("username", "alert")
		// 接受其他的
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 多选框
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v",
			type1, username, password, hobbys))
	})
	r.PUT("/xxxPut")

	// 限制表单上传大小 8MB 默认为32MB
	//r.MaxMultipartMemory - 8 << 20
	// 上传单个图片
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 传到项目的更目录， 名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		// 打印信息
		c.String(http.StatusOK, fmt.Sprintf("%s upload!", file.Filename))
	})

	// 上传多个图片
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历上述所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// 路由组1，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	getJson(r)
	loginForm(r)
	loginUri(r)

	// 3、监听端口
	r.Run(":8111")
}

func getting(c *gin.Context) {

}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "hellow %s\n", name)
}

func submit(c *gin.Context) {
	fmt.Println("submit")
}

// json获取数据
func getJson(r *gin.Engine) {
	r.POST("loginJson", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// 表单获取数据
func loginForm(r *gin.Engine) {
	r.POST("loginForm", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind() 默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// uri数据解析和绑定
func loginUri(r *gin.Engine) {
	r.GET("loginUri/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

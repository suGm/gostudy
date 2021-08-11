package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sgm/web/service"
	"net/http"
	"strconv"
)

// 访问主页的控制器
func IndexHandle(c *gin.Context) {
	// 从Service取数据
	// 1、加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 2、加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	// gin.H本质上是个map
	//var data map[string]interface{} = make(map[string]interface{}, 16)
	//data["article_list"] = articleRecordList
	//data["category_list"] = categoryList
	//
	//c.HTML(http.StatusOK, "views/index.html", data)

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	// 转
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 根据分类id再次获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

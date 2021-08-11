package db

import (
	"github.com/sgm/web/model"
	"testing"
	"time"
)

func init() {
	// parseTime = true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加则报错
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 测试插入文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "abc"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "123"
	article.ArticleInfo.Username = "sgm"
	article.ArticleInfo.Summary = "abc"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId: %d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		return
	}
	t.Logf("article:%d\n", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(1)
	if err != nil {
		return
	}
	t.Logf("article:%#v\n", detail.Username)
}

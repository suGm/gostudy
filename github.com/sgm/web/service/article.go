package service

import (
	"github.com/sgm/web/dao/db"
	"github.com/sgm/web/model"
)

// 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2。获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		// 根据当前端的文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章取出分类id
		categoryId := article.CategoryId
		// 遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据分类id，获取该类文章和他们对应分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2。获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		// 根据当前端的文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章取出分类id
		categoryId := article.CategoryId
		// 遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据多个文章的id,获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	// 遍历文章，得到每个文章
	for _, article := range articleInfoList {
		flag := true
		// 从当前文章去除分类id
		categegoryId := article.CategoryId
		// 去重
		for _, id := range ids {
			// 看当前id是否存在
			if id == categegoryId {
				flag = false
				break
			}
		}
		if flag {
			ids = append(ids, categegoryId)
		}
	}
	return
}

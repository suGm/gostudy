package service

import (
	"github.com/sgm/web/dao/db"
	"github.com/sgm/web/model"
)

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllGategoryList()
	if err != nil {
		return
	}
	return
}

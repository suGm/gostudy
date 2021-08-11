package db

import "testing"

func init() {
	// parseTime = true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加则报错
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	list, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		t.Logf("id:%d category:%#v", v.CategoryId, v)
	}
}

func TestGetAllGategoryList(t *testing.T) {
	list, err := GetAllGategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		t.Logf("id:%d category:%#v", v.CategoryId, v)
	}
}

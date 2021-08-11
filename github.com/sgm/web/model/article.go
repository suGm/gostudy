package model

import "time"

// 定义文章结构体
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"` // 文章摘要
	Title        string    `db:"title"`
	CreateTime   time.Time `db:"create_time"` // 时间
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// 为了提升效率
type ArticleDetail struct {
	ArticleInfo
	// 文章内容
	Content string `db:"content"`
	Category
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}

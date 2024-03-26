package models

// 从数据库获取首页的分类模板
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

// 设置分类页面的数据模板
type CategoryResponse struct {
	*HomeResoponse
	CategoryName string
}

package service

import (
	"html/template"
	"my-project/config"
	"my-project/dao"
	"my-project/models"
)

//用于index页面的sql操作

func GetAllIndexInfo(page, pageSize int) (*models.HomeResoponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	posts, err := dao.GetPostpage(page, pageSize)
	// 由于执行模板中Posts为postmore 需要根据post中id获取postmore的内容
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		//转为中文字符 判断只显示前100个字
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//获取文章总数
	total := dao.CountGetAllPost()
	pagesCount := (total-1)/10 + 1
	var pages []int
	// 将页码存储在数组中
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	// 执行模板
	var hr = &models.HomeResoponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	return hr, nil
}

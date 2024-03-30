package service

import (
	"html/template"
	"my-project/config"
	"my-project/dao"
	"my-project/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	if err != nil {
		return nil, err
	}
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	postMore.Pid = post.Pid
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}

package dao

import (
	"log"
	"my-project/models"
)

// 总页面数据
func GetPostpage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostpage 取值失败")
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil

}

// 根据分类获取页面数据
func GetPostpageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostpage 取值失败")
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil

}

// 获取总页码数
func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post;")
	_ = rows.Scan(&count)
	return
}

// 根据分类获取页码数
func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?;", cId)
	_ = rows.Scan(&count)
	return
}

// 获取文章详情
func GetPostById(pid int) (models.Post, error) {
	rows := DB.QueryRow("select * from blog_post where pid = ?;", pid)
	var post models.Post
	if rows.Err() != nil {
		return post, rows.Err()
	}

	err := rows.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, err
	}
	return post, nil
}

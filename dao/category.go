package dao

import (
	"log"
	"my-project/models"
)

// 根据cid获取分类名
func GetCategoryNameById(Cid int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", Cid)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

// 获取分类全部信息
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询失败")
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值失败")
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

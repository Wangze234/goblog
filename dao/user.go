package dao

import (
	"log"
	"my-project/models"
)

// 根据cid获取用户名
func GetUserNameById(Uid int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", Uid)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var UserName string
	_ = row.Scan(&UserName)
	return UserName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd= ? limit 1",
		userName,
		passwd,
	)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

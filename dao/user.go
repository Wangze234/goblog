package dao

import "log"

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

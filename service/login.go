package service

import (
	"errors"
	"my-project/dao"
	"my-project/models"
)

func Login(userName, passwd string) (*models.LoginRes, error) {

	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	var lr = &models.LoginRes{}
	return lr, nil
}

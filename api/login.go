package api

import (
	"my-project/common"
	"my-project/service"
	"net/http"
)

func (*APIhander) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码 返回对应的json数据
	//返回json数据
	param := common.GetRequestJsonParam(r)
	userName := param["username"].(string)
	passwd := param["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	//返回成功或失败数据
	common.Success(w, loginRes)
}

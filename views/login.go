package views

import (
	"my-project/common"
	"my-project/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	//获取index模板
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}

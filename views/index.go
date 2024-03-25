package views

import (
	"errors"
	"log"
	"my-project/common"
	"my-project/service"
	"net/http"
)

type IndexData struct {
	Tile string `json:"tile"`
	Desc string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// 页面上涉及到的数据必须有定义
	//获取index模板
	index := common.Template.Index
	// 数据库查询
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误， 请联系管理员！"))
	}

	//向index模板传值
	index.WriteData(w, hr)
}

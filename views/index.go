package views

import (
	"errors"
	"log"
	"my-project/common"
	"my-project/service"
	"net/http"
	"strconv"
)

type IndexData struct {
	Tile string `json:"tile"`
	Desc string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// 页面上涉及到的数据必须有定义
	//获取index模板
	index := common.Template.Index
	// 获取表单
	if err := r.ParseForm(); err != nil {
		log.Println("获取表单失败：", err)
		index.WriteError(w, errors.New("系统错误， 请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	// 数据库查询
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误， 请联系管理员！"))
	}

	//向index模板传值
	index.WriteData(w, hr)
}

package views

import (
	"my-project/common"
	"my-project/config"
	"my-project/models"
	"net/http"
)

type IndexData struct {
	Tile string `json:"tile"`
	Desc string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// 页面上涉及到的数据必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	// 执行模板
	var hr = &models.HomeResoponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	//获取index模板
	index := common.Template.Index
	//向index模板传值
	index.WriteData(w, hr)
}

package router

import (
	"my-project/api"
	"my-project/views"
	"net/http"
)

func Router() {
	// 路由分三类 1.页面  2.api 数据（json） 3.静态资源
	//页面
	http.HandleFunc("/", views.HTML.Index)
	//api接口 数据
	http.HandleFunc("/post", api.API.SaveUpdatePost)

	//映射resource 静态资源的路径
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
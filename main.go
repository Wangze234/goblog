package main //声明文件所在的包 每一个文件都必须归属一个包

import (
	"fmt"
	"log"
	"my-project/common"
	"my-project/router"
	"net/http"
) //引入程序中使用的包 使用其中的函数

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	fmt.Println("Hello World")
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	//路由功能
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

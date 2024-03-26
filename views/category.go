package views

import (
	"errors"
	"log"
	"my-project/common"
	"my-project/service"
	"net/http"
	"strconv"
	"strings"
)

// 分类页监听信息 获取分类cid
func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	// 拿到分类的页面 在页面中返回错误信息
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此路径"))
		return
	}

	//设置分类页面的分页
	// 获取表单
	if err := r.ParseForm(); err != nil {
		log.Println("获取表单失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误， 请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)

}

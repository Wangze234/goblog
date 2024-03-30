package views

import (
	"errors"
	"my-project/common"
	"my-project/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	// 拿到分类的页面 在页面中返回错误信息
	detail := common.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	// 得到x.html 需要删除html后缀
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	// 获取pid
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此路径"))
		return
	}
	//查询数据库
	// 最终需要返回的值是
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)

}

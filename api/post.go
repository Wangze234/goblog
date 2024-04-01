package api

import (
	"errors"
	"my-project/common"
	"my-project/models"
	"my-project/service"
	"my-project/utils"
	"net/http"
	"strconv"
	"time"
)

func (*APIhander) SaveUpdatePost(w http.ResponseWriter, r *http.Request) {
	//判断用户是否登录 鉴权
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)

	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST表示save操作
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pintType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pintType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)

	}
}

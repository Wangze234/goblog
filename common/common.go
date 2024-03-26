package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"my-project/config"
	"my-project/models"
	"net/http"
	"sync"
)

// 给对应的路径赋值
var Template models.HtmlTemplate

// 给Template赋值
func LoadTemplate() {
	// 增加一个协程
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		// 耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

// 获取json参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	//读取body
	body, _ := ioutil.ReadAll(r.Body)
	//设置map存储
	_ = json.Unmarshal(body, &params)
	return params
}

// 登录失败 返回状态码和数据
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// 登录成功 返回状态码和数据
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

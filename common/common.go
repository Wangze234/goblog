package common

import (
	"my-project/config"
	"my-project/models"
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

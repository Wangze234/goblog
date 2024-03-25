package common

import (
	"my-project/config"
	"my-project/models"
)

// 给对应的路径赋值
var Template models.HtmlTemplate

// 给Template赋值
func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}

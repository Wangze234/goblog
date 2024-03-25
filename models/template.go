package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

// 判断导航栏标签位置
func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

//解析模板的匹配，index只需负责传输数据，模板匹配任务由template完成

// 自己定义的template模板类型，可以加入一些数据操作函数，再封装，以便传值
type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index     TemplateBlog
	Category  TemplateBlog
	Custom    TemplateBlog
	Detail    TemplateBlog
	Login     TemplateBlog
	Pigenhole TemplateBlog
	Writing   TemplateBlog
}

// 模板执行excotor函数
func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte("error"))
	}
}

// 入参 模板路径
func InitTemplate(templateDir string) HtmlTemplate {
	//期望读取路径后的返回
	tp := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigenhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate
}

// 入参 模板的组合 希望返回TemplateBlog类型的数组
func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	//templates是拿到的页面名称
	//templateDir是当前html页面所在路径
	for _, view := range templates {
		viewName := view + ".html"
		//设置路径与文件加载
		// new一个解析模板，命名为index.html
		t := template.New(viewName)
		//访问博客首页模板时，因为有多个模板嵌套，解析文件的时候需要将所有涉及的模板进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println(err)
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs
}

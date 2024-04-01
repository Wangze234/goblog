package views

import (
	"my-project/common"
	"my-project/service"
	"net/http"
)

func (a *HTMLApi) Write(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}

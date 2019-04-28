package user

import (
	"petshop/controllers"
	"petshop/models"
)

type UnRegisterController struct {
	controllers.BaseController
}

func (this *UnRegisterController) Get() {
	this.Verify(this.User())
	this.TplName = "unregister.tpl"
}

func (this *UnRegisterController) Post() {
	this.Verify(this.User())
	username := this.GetSession("username")
	if username != nil {
		err := models.UnRegister(username.(string))
		if err == nil {
			this.Redirect("/", 302)
			return
		}
	}
	this.TplName = "error.tpl"
	return
}

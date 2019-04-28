package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type DelUserController struct {
	controllers.BaseController
}

func (this *DelUserController) Get() {
	this.Verify(this.Admin())
	this.TplName = "delUser.tpl"
}

func (this *DelUserController) Post() {
	this.Verify(this.Admin())
	username := this.GetString("username")
	err := models.UnRegister(username)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.TplName = "success.tpl"
	return
}

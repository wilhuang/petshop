package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type GetUserlistController struct {
	controllers.BaseController
}

func (this *GetUserlistController) Get() {
	this.Verify(this.Admin())
	this.TplName = "login.tpl"
}

func (this *GetUserlistController) Post() {
	this.Verify(this.Admin())
	list, err := models.GetAllUser()
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Data["list"] = list
	this.TplName = "userlist.tpl"
	return
}

package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type AddUserController struct {
	controllers.BaseController
}

func (this *AddUserController) Get() {
	this.Verify(this.Admin())
	this.TplName = "addUser.tpl"
}

func (this *AddUserController) Post() {
	this.Verify(this.Admin())
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	err := models.Register(username, password, email)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.TplName = "success.tpl"
	return
}

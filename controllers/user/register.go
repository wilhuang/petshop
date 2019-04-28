package user

import (
	"petshop/controllers"
	"petshop/models"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) Get() {
	this.TplName = "register.tpl"
}

func (this *RegisterController) Post() {
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

package common

import (
	"petshop/controllers"
	"petshop/models"

	"github.com/astaxie/beego"
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
	this.Data["Finish"] = true
	if username == "" || password == "" || email == "" {
		this.TplName = "register.tpl"
		return
	}
	err := models.Register(username, password, email)
	beego.Info(err)
	if err == nil {
		this.Data["UserName"] = username
	}
	this.TplName = "register.tpl"
	return
}

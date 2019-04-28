package user

import (
	"petshop/controllers"
	"petshop/models"
)

type ResetPwdController struct {
	controllers.BaseController
}

func (this *ResetPwdController) Get() {
	this.Verify(this.User())
	this.TplName = "resetpwd.tpl"
}

func (this *ResetPwdController) Post() {
	this.Verify(this.User())
	password := this.GetString("password")
	username, ok := this.GetSession("username").(string)
	if !ok {
		this.TplName = "error.tpl"
		return
	}
	err := models.ResetPwd(username, password)
	if err != nil {
		this.TplName = "resetpwd.tpl"
		return
	}
	this.Redirect("/user", 302)
	return
}

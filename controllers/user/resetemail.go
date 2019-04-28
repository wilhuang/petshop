package user

import (
	"petshop/controllers"
	"petshop/models"
)

type ResetEmailController struct {
	controllers.BaseController
}

func (this *ResetEmailController) Get() {
	this.Verify(this.User())
	this.TplName = "resetemail.tpl"
}

func (this *ResetEmailController) Post() {
	this.Verify(this.User())
	email := this.GetString("email")
	username, ok := this.GetSession("username").(string)
	if !ok {
		this.TplName = "error.tpl"
		return
	}
	err := models.ResetEmail(username, email)
	if err != nil {
		this.TplName = "resetemail.tpl"
		return
	}
	this.Redirect("/user", 302)
	return
}

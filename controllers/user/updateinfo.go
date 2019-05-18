package user

import (
	"petshop/controllers"
	"petshop/models"
)

type UpdateInfoController struct {
	controllers.BaseController
}

func (this *UpdateInfoController) Post() {
	this.Verify(this.User())
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	this.Data["Finish"] = true
	if username == "" || password == "" || email == "" {
		this.SsionToTpl()
		this.TplName = "user.tpl"
		return
	}
	uid, ok := this.GetSession("Uid").(string)
	if !ok {
		this.TplName = "error.tpl"
		return
	}
	err := models.UpdateInfo(uid, username, password, email)
	if err == nil {
		this.DelSession("UserName")
		this.DelSession("Email")
		this.SetSession("UserName", username)
		this.SetSession("Email", email)
		this.Data["OK"] = true
	}
	this.SsionToTpl()
	this.TplName = "user.tpl"
	return
}

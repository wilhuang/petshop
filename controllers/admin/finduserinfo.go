package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type FindUserInfoController struct {
	controllers.BaseController
}

func (this *FindUserInfoController) Get() {
	this.Verify(this.Admin())
	this.TplName = "findUserInfo.tpl"
}

func (this *FindUserInfoController) Post() {
	this.Verify(this.Admin())
	username := this.GetString("username")
	info, err := models.FindUserByUserName(username)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Data["Uid"] = info.Uid
	this.Data["UserName"] = info.UserName
	this.Data["Role"] = info.Role
	this.Data["Email"] = info.Email
	this.TplName = "userinfo.tpl"
	return
}

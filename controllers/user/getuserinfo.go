package user

import (
	"petshop/controllers"
)

type GetUserInfoController struct {
	controllers.BaseController
}

func (this *GetUserInfoController) Get() {
	this.Verify(this.User())
	this.SsionToTpl()
	this.TplName = "userinfo.tpl"
	return
}

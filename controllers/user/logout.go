package user

import (
	"petshop/controllers"
)

type LogoutController struct {
	controllers.BaseController
}

func (this *LogoutController) Get() {
	this.Verify(this.User())
	u := this.GetSession("UserName")
	if u != nil {
		this.DestroySession()
	}
	this.TplName = "login.tpl"
	return
}

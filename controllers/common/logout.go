package common

import (
	"petshop/controllers"
)

type LogoutController struct {
	controllers.BaseController
}

func (this *LogoutController) Get() {
	u := this.GetSession("UserName")
	if u != nil {
		this.DestroySession()
	}
	this.Redirect("/login", 302)
	return
}

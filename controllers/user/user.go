package user

import "petshop/controllers"

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Get() {
	this.Verify(this.User())
	this.SsionToTpl()
	this.TplName = "index.tpl"
	return
}

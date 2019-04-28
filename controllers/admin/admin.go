package admin

import "petshop/controllers"

type AdminController struct {
	controllers.BaseController
}

func (this *AdminController) Get() {
	this.Verify(this.Admin())
	this.SsionToTpl()
	this.TplName = "admin.tpl"
	return
}

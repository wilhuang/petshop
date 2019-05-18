package common

import "petshop/controllers"

type RoleController struct {
	controllers.BaseController
}

func (this *RoleController) Get() {
	this.TplName = "role.tpl"
	return
}

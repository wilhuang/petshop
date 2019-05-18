package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type DelProductController struct {
	controllers.BaseController
}

func (this *DelProductController) Get() {
	this.Verify(this.Admin())
	pid := this.Ctx.Input.Param(":id")
	err := models.UnRegisterById(pid)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.TplName = "success.tpl"
	return
}

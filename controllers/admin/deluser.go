package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type DelUserController struct {
	controllers.BaseController
}

func (this *DelUserController) Get() {
	this.Verify(this.Admin())
	uid := this.Ctx.Input.Param(":id")
	err := models.UnRegisterById(uid)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.TplName = "success.tpl"
	return
}

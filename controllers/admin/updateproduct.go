package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type UpdateProductController struct {
	controllers.BaseController
}

func (this *UpdateProductController) Get() {
	this.Verify(this.Admin())
	this.TplName = "findUserInfo.tpl"
}

func (this *UpdateProductController) Post() {
	this.Verify(this.Admin())
	pid := this.Ctx.Input.Param(":id")
	pnmae := this.GetString("pname")
	oprice := this.GetString("oprice")
	onhand := this.GetString("onhand")
	if pnmae == "" || oprice == "" || onhand == "" || pid == "" {
		this.Redirect("/admin", 302)
		return
	}
	models.UpdateDetails(pid, pnmae, oprice, onhand)
	this.Redirect("/admin", 302)
	return
}

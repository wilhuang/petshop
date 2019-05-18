package cart

import (
	"petshop/controllers"
	"petshop/models"

	"github.com/astaxie/beego"
)

type AddController struct {
	controllers.BaseController
}

func (this *AddController) Get() {
	this.Verify(this.User())
	pid := this.Ctx.Input.Param(":pid")
	uid := this.GetSession("Uid").(string)
	err := models.AddCart(uid, pid)
	if err != nil {
		beego.Debug(err)
		this.TplName = "error.tpl"
		return
	}
	this.Redirect("/cart", 302)
	return
}

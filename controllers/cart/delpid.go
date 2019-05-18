package cart

import (
	"petshop/controllers"
	"petshop/models"

	"github.com/astaxie/beego"
)

type DelpidController struct {
	controllers.BaseController
}

func (this *DelpidController) Get() {
	this.Verify(this.User())
	uid := this.GetSession("Uid").(string)
	pid := this.Ctx.Input.Param(":pid")
	beego.Debug(uid, pid)
	err := models.DeleteCart(uid, pid)
	beego.Debug(err)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Redirect("/cart", 302)
	return
}

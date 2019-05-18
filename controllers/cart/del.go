package cart

import (
	"petshop/controllers"
	"petshop/models"
	"strconv"

	"github.com/astaxie/beego"
)

type DelController struct {
	controllers.BaseController
}

func (this *DelController) Get() {
	this.Verify(this.User())
	pid := this.Ctx.Input.Param(":pid")
	uid := this.GetSession("Uid").(string)
	info, err := models.CheckCart(uid, pid)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	beego.Debug(pid, info.Pnum)
	num, err := strconv.Atoi(info.Pnum)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	if num < 2 {
		err = models.DeleteCart(uid, pid)
		if err != nil {
			this.TplName = "error.tpl"
			return
		}
		this.Redirect("/cart", 302)
		return
	}
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	err = models.UpdatePnum(uid, pid, strconv.FormatInt(int64(num-1), 10))
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Redirect("/cart", 302)
	return
}

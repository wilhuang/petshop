package cart

import (
	"petshop/controllers"
	"petshop/models"

	"github.com/astaxie/beego"
)

type CartInfo struct {
	Pid    string
	Pname  string
	Oprice string
	Pimg1  string
	Pnum   string
}

type CartController struct {
	controllers.BaseController
}

func (this *CartController) Get() {
	this.Verify(this.User())
	uid, _ := this.GetSession("Uid").(string)
	beego.Info(uid)
	allnum, allprice, list, err := models.CartShow(uid)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	beego.Info(list)
	this.Data["AllNum"] = allnum
	this.Data["AllPrice"] = allprice
	this.Data["CartList"] = list
	this.TplName = "cart.tpl"
	return
}

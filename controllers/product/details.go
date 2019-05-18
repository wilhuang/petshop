package product

import (
	"petshop/controllers"
	"petshop/models"

	"github.com/astaxie/beego"
)

type DetailsController struct {
	controllers.BaseController
}

func (this *DetailsController) Get() {
	pid := this.Ctx.Input.Param(":pid")
	beego.Info(pid)
	info, err := models.FindProductByPid(pid)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	beego.Info(info)
	this.Data["Pid"] = info.Pname
	this.Data["Pname"] = info.Pname
	this.Data["Oprice"] = info.Oprice
	this.Data["Onhand"] = info.Onhand
	this.Data["Pimg1"] = info.Pimg1
	this.Data["Pimg2"] = info.Pimg2
	this.Data["Pimg3"] = info.Pimg3
	this.Data["Pimg4"] = info.Pimg4
	this.Data["Pimg5"] = info.Pimg5
	this.Data["Pimg6"] = info.Pimg6
	this.Data["Pnamea"] = "猫粮"
	this.Data["Pnameb"] = "狗粮"
	this.Data["Pimga"] = "../static/img/product/008B.jpg"
	this.Data["Pimgb"] = "../static/img/product/018A.jpg"
	this.Data["Opricea"] = "119"
	this.Data["Opriceb"] = "169"
	this.TplName = "details.tpl"
	return
}

package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type AdminController struct {
	controllers.BaseController
}

func (this *AdminController) Get() {
	this.Verify(this.Admin())
	this.SsionToTpl()
	usernum, list1, err := models.GetAllUser()
	productnum, list2, err := models.GetAllProduct()
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Data["UserList"] = list1
	this.Data["usernum"] = usernum
	this.Data["ProductList"] = list2
	this.Data["productnum"] = productnum
	this.TplName = "admin.tpl"
	return
}

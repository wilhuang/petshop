package common

import (
	"petshop/controllers"
	"petshop/models"
)

type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
	list1, list2, list3, err := models.GetProductSplit()
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	this.Data["ProductList1"] = list1
	this.Data["ProductList2"] = list2
	this.Data["ProductList3"] = list3
	username := this.GetSession("UserName")
	if username != nil {
		this.Data["UserName"] = this.GetSession("UserName")
	}
	this.TplName = "index.tpl"
	return
}

func (this *MainController) Post() {
	username := this.GetSession("UserName")
	if username != nil {
		this.Data["UserName"] = this.GetSession("UserName")
	}
	this.TplName = "index.tpl"
	return
}

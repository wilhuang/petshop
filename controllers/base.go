package controllers

import (
	"github.com/astaxie/beego"
)

type Role int

// //equal Role.Visitor()=0
// func (r *Role) Visitor() int {
// 	return 0
// }

//equal Role.User()=1
func (r *Role) User() int {
	return 0
}

//equal Role.Admin()=2
func (r *Role) Admin() int {
	return 1
}

type BaseController struct {
	beego.Controller
	Role
}

func (this *BaseController) Verify(powerAccessPage int) {
	role, _ := this.GetSession("Role").(int)
	username := this.GetSession("UserName")
	if username == nil || role+1 < powerAccessPage {
		this.Redirect("/role", 302)
		return
	}
}

//send which get from session to tpl
func (this *BaseController) SsionToTpl() {
	this.Data["Time"] = this.GetSession("Time")
	this.Data["Uid"] = this.GetSession("Uid")
	this.Data["UserName"] = this.GetSession("UserName")
	this.Data["Role"] = this.GetSession("Role")
	this.Data["Email"] = this.GetSession("Email")
}

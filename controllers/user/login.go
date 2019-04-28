package user

import (
	"petshop/controllers"
	"petshop/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	info, err := models.Login(username, password)
	beego.Info(info)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	beego.Info(info.UserName)
	this.SetSession("Uid", info.Uid)
	this.SetSession("UserName", info.UserName)
	this.SetSession("Role", info.Role)
	this.SetSession("Email", info.Email)
	this.SetSession("Time", time.Now().Format("15:04:05"))
	role, err := strconv.Atoi(info.Role)
	if err != nil {
		this.TplName = "login.tpl"
		return
	}
	// role, _ := info.Role.(int)
	beego.Info(role)
	// if !ok {
	// 	this.TplName = "login.tpl"
	// 	return
	// }
	if role == this.Admin() {
		this.Redirect("/admin", 302)
		return
	}
	this.Redirect("/user", 302)
	return
}

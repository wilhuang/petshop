package common

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
	beego.Info(username, password)
	info, err := models.Login(username, password)
	if err != nil {
		this.Redirect("/login", 302)
		return
	}
	this.SetSession("Uid", info.Uid)
	this.SetSession("UserName", info.UserName)
	this.SetSession("Role", info.Role)
	this.SetSession("Email", info.Email)
	this.SetSession("Time", time.Now().Format("15:04:05"))
	role, err := strconv.Atoi(info.Role)
	if err != nil {
		this.TplName = "error.tpl"
		return
	}
	if role == this.Admin() {
		this.Redirect("/admin", 302)
		return
	}
	this.Redirect("/", 302)
	return
}

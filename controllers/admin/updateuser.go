package admin

import (
	"petshop/controllers"
	"petshop/models"
)

type UpdateUserController struct {
	controllers.BaseController
}

func (this *UpdateUserController) Post() {
	this.Verify(this.Admin())
	uid := this.Ctx.Input.Param(":id")
	username := this.GetString("username")
	email := this.GetString("email")
	if username == "" || uid == "" || email == "" {
		this.Redirect("/admin", 302)
		return
	}
	models.UpdateInfoAdmin(uid, username, email)
	this.Redirect("/admin", 302)
	return
}

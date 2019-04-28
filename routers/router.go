package routers

import (
	"petshop/controllers/admin"
	"petshop/controllers/commodity"
	"petshop/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &user.LoginController{})
	beego.Router("/user", &user.UserController{})
	beego.Router("/register", &user.RegisterController{})
	beego.Router("/unregister", &user.UnRegisterController{})
	beego.Router("/logout", &user.LogoutController{})
	beego.Router("/userinfo", &user.GetUserInfoController{})
	beego.Router("/userinfo/resetemail", &user.ResetEmailController{})
	beego.Router("/userinfo/resetpwd", &user.ResetPwdController{})
	beego.Router("/admin", &admin.AdminController{})
	beego.Router("/admin/getuserlist", &admin.GetUserlistController{})
	beego.Router("/admin/finduserinfo", &admin.FindUserInfoController{})
	beego.Router("/admin/adduser", &admin.AddUserController{})
	beego.Router("/admin/deluser", &admin.DelUserController{})
	beego.Router("/commodity", &commodity.CommodityController{})
}

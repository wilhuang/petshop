package routers

import (
	"petshop/controllers/admin"
	"petshop/controllers/cart"
	"petshop/controllers/common"
	"petshop/controllers/product"
	"petshop/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &common.MainController{})
	beego.Router("/register", &common.RegisterController{})
	beego.Router("/login", &common.LoginController{})
	beego.Router("/logout", &common.LogoutController{})
	beego.Router("/role", &common.RoleController{})

	beego.Router("/user", &user.UserController{})
	beego.Router("/user/updateinfo", &user.UpdateInfoController{})

	beego.Router("/product/?:pid", &product.DetailsController{})

	beego.Router("/cart", &cart.CartController{})
	beego.Router("/cart/del/?:pid", &cart.DelController{})
	beego.Router("/cart/add/?:pid", &cart.AddController{})
	beego.Router("/cart/delp/?:pid", &cart.DelpidController{})

	beego.Router("/admin", &admin.AdminController{})
	beego.Router("/admin/deluser/?:id", &admin.DelUserController{})
	beego.Router("/admin/updateuser/?:id", &admin.UpdateUserController{})
	beego.Router("/admin/updateproduct/?:id", &admin.UpdateUserController{})
	beego.Router("/admin/delproduct/?:id", &admin.DelUserController{})
	// beego.Router("/admin/finduserinfo", &admin.FindUserInfoController{})
}

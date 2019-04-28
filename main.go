package main

import (
	_ "petshop/routers"

	"github.com/astaxie/beego"
)

func main() {
	// beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

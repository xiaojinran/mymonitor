package controllers

import (
	"github.com/astaxie/beego"
)

type HomeContrller struct {
	beego.Controller
}

func (this *HomeContrller) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "index.html"
	this.Data["user"] = this.Input().Get("username")

}

package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {


	if beego.AppConfig.String("uname") == this.Input().Get("uname") {
		this.Ctx.Redirect(302, "/")
		return

	} else {
		this.Ctx.Redirect(302, "/logout")
		return
	}

}

package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController)Get(){
	this.TplName = "register.html"
}

func (this *RegisterController)Post(){
	var user models.User
	user.Name = this.Input().Get("uname")
	user.Pass =  this.Input().Get("passwd")
	models.AddUser(&user)
	this.Ctx.Redirect(302,"/login")
	return
}
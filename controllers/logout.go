package controllers

import (
	"github.com/astaxie/beego"
)


type LogOutController struct{
	beego.Controller
}

func (this *LogOutController)Get(){
	this.Ctx.Redirect(302,"/login")
	return
}
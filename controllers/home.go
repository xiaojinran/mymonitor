package controllers

import (
	"github.com/astaxie/beego"
)

type HomeContrller struct {
	beego.Controller
}

func (this *HomeContrller) Get() {
	this.TplName = "index.html"

}

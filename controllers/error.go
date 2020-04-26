package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

//404页面处理
func (this *ErrorController) Error404() {
	this.Layout="layout.tpl"
	this.TplName="404.html"
}
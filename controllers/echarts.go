package controllers

import "github.com/astaxie/beego"

type EchartsController struct {
	beego.Controller
}


func (this *EchartsController)Get(){
	this.Layout="layout.tpl"
	this.TplName="echarts.html"
}
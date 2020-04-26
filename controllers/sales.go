package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type SalesController struct {
	beego.Controller
}

// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /sales [Get]
func (c *SalesController) GetSales(){
	catalog := []string{"a","b","c"}
	count := []int{33,44,55}
	ret:=models.NewSales(catalog,count)
	c.Data["json"] = ret
	c.ServeJSON()
}

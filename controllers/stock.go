package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type StockController struct {
	beego.Controller
}
var stockList []*models.Stock

func (c *StockController)Get(){
	c.Layout = "layout.tpl"
	c.TplName = "stock.html"
	c.Data["alert"] = alert
    c.Data["alertmsg"] =alertmsg
	c.Data["stock"] = stockList
}

func (c *StockController)Post(){
	stockid:=c.Input().Get("stockid")
    if stockid == ""{
	    c.Redirect("/stock",302)
	    return
    }
	s1,err:=models.GetStockMsg(stockid)
	if err !=nil{
		alert= true
		alertmsg = err.Error()
		c.Redirect("/stock",302)
		return
	}
	stockList =append(stockList, s1)
	alert= false
	c.Redirect("/stock",302)
	return
	
	
}
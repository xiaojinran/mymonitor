package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type ProductController struct{
	beego.Controller
}

// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [Get]
func (c *ProductController)GetProduct(){
	if c.Input().Get("delid")!=""{
		models.DeleteProduct(&models.Product{
			Id:StrParseInt(c.Input().Get("delid")),
		})
	}
	c.Layout = "layout.tpl"
	c.TplName = "products.html"
	c.Data["products"] = models.GetAllProduct()
}

// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /add [Post]
func (c *ProductController)AddProduct(){
	name := c.Input().Get("name")
	platform := c.Input().Get("platform")
	_type := c.Input().Get("type")
	url := c.Input().Get("url")
	price := c.Input().Get("price")
	_price := StrParseFloat(price)
	product := models.NewProduct(name,_type,platform,url,_price)
	models.AddProduct(product)
	c.Redirect("/v1/product",302)
	return
}

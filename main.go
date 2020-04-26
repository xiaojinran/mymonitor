package main

import (
	"github.com/astaxie/beego"
	"mymonitor/controllers"
	_ "mymonitor/models"
	_ "mymonitor/routers"
	_ "mymonitor/models"
)

func main() {
	beego.SetStaticPath("/assets","assets")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}


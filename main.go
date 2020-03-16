package main

import (
	"github.com/astaxie/beego"
	_ "mymonitor/models"
	_ "mymonitor/routers"
	_ "mymonitor/models"
)

func main() {
	beego.SetStaticPath("/assets","assets")
	beego.Run()
}


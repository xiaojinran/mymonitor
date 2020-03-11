package main

import (
	_ "mymonitor/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/assets","assets")
	beego.Run()
}


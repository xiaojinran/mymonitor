package routers

import (
	"mymonitor/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeContrller{})
	beego.Router("/logout", &controllers.LogOutController{})
	beego.Router("/login", &controllers.LoginController{})

}

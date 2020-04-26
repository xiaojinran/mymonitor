package routers

import (
	"mymonitor/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeContrller{})
	
	beego.Router("/logout", &controllers.LogOutController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/wechat",&controllers.WechatController{})
	beego.Router("/echarts",&controllers.EchartsController{})
	beego.Router("/stock",&controllers.StockController{})
	
	ns := beego.NewNamespace("/config",
		beego.NSInclude(&controllers.ConfigController{},
		),

	)
	
	
	api := beego.NewNamespace("/v1",
		beego.NSNamespace("/qywechat",
			beego.NSInclude(
				&controllers.WechatController{},
			),
	),
		beego.NSNamespace("/data",
			beego.NSInclude(
				&controllers.SalesController{},
			),
			
		),
		beego.NSNamespace("/linux",
			beego.NSInclude(
				&controllers.LinuxController{},
			),
		
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		
		),
		
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(api)

}

package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
           this.TplName="login.html"
}

func (this *LoginController) Post() {
	//通过sqlite数据库验证账号密码是否正确
     user := models.GetUserByName(this.Input().Get("uname"))
     if user.Pass == this.Input().Get("passwd") {
	     this.Ctx.Redirect(302, "/")
	     return
	
     }else {
	     this.Ctx.Redirect(302, "/logout")
	     return
     }

}

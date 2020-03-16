package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mymonitor/models"
	"mymonitor/utils"
	"strconv"
	"time"
)

type WechatController struct{
	beego.Controller
}

func (this *WechatController)Get (){
	delid := this.Input().Get("delid")
	id , _ := strconv.ParseInt(delid, 10, 32)
	if delid != "" {
		msg:= &models.Message{
		   Id: int(id),
		}
		models.DelMessage(msg)
		this.Redirect("/wechat",302)
		return
	}else{
		userlist := utils.GetAllUser(GetToken())
		messages := models.GetAllMessage()
		fmt.Println(messages)
		this.TplName="wechat.html"
		fmt.Println(userlist.UserList)
		this.Data["userlist"] = userlist.UserList
		this.Data["messages"] = messages
		//this.Redirect("/wechat",302)
	}

	
}

func (this *WechatController)Post(){
	touser:=this.Input().Get("touser")
	message := this.Input().Get("message")
	utils.SendToUser(touser,message,GetToken())
	msg := &models.Message{
		User:touser,
		Msg:message,
		Time:time.Now(),
		
	}
	models.AddMessage(msg)
	this.Redirect("/wechat",302)
	return
}
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"mymonitor/models"
	"strconv"
	"time"
)

type WechatController struct{
	beego.Controller
}

func (c *WechatController)Get (){
	delid := c.Input().Get("delid")
	id , _ := strconv.ParseInt(delid, 10, 32)
	if delid != "" {
		msg:= &models.Message{
		   Id: int(id),
		}
		models.DelMessage(msg)
		c.Redirect("/wechat",302)
		return
	}else{
		userlist := models.GetAllWXUser(GetToken())
		messages := models.GetAllMessage()
		c.Layout = "layout.tpl"
		c.TplName="wechat.html"
		fmt.Println(userlist.UserList)
		c.Data["userlist"] = userlist.UserList
		c.Data["messages"] = messages
	}

	
}

func (c *WechatController)Post(){
	touser:= c.Input().Get("touser")
	message := c.Input().Get("message")
	models.SendToUser(touser,message,GetToken())
	msg := &models.Message{
		User:touser,
		Msg:message,
		Time:time.Now(),
		
	}
	models.AddMessage(msg)
	c.Redirect("/wechat",302)
	return
}


// @Title Post
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /send [Post]
func (c *WechatController) WechatAPI(){
	message := c.Ctx.Input.RequestBody
	fmt.Println("fffff:" + string(message))
	var msg models.Message
	json.Unmarshal(message,&msg)
	fmt.Println(msg.Msg)
	ret := models.SendToUser(msg.User, msg.Msg, GetToken())
	
	c.Data["json"] = ret
	c.ServeJSON()
}

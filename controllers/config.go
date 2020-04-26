package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
)

type ConfigController struct {
	beego.Controller
}

// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /ssh [Get]
func (c *ConfigController) SshConfig(){
	delid := c.Input().Get("delid")
	if delid !=""{
	  id := StrParseInt(delid)
	  ssh := &models.SshHost{
	  	Id:id,
	  }
		models.DeleteSshHost(ssh)
	  c.Redirect("/config/ssh",302)
	}
	c.Layout = "layout.tpl"
	c.TplName = "ssh_conf.html"
	c.Data["alert"] = alert
	c.Data["alertmsg"] =alertmsg
	c.Data["sshhost"] = models.GetAllSshHost()
	

	
}

// @Title Post
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /ssh [Post]
func (c *ConfigController) AddSshHost(){
	host := c.Input().Get("host")
	user:=c.Input().Get("user")
	password := c.Input().Get("password")
	port := c.Input().Get("port")
	port_int := StrParseInt(port)
	sshHost := models.NewSshHost(host, user, password, port_int)
	e := models.AddSSHHost(sshHost)
	if e!=nil{
	    alert=true
	    alertmsg=e.Error()
	}
	c.Redirect("/config/ssh",302)
	
}


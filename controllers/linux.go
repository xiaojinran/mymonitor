package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mymonitor/models"
	"strings"
)

type LinuxController struct {
	beego.Controller
}
// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /details [Get]
func (c *LinuxController)LinuxDetails(){
	c.Layout = "layout.tpl"
	c.TplName = "linux_details.html"
}
// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /ops [Get]
func (c *LinuxController)LinuxOps(){
	c.Layout = "layout.tpl"
	c.TplName = "linux_ops.html"
	i := c.Input().Get("id")
	id := StrParseInt(i)
	host := models.GetSshHost(id)
	c.Data["result"] = models.GetAllSshResult()
	c.Data["host"] = host
}

// @Title Post
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /ops [Post]
func (c *LinuxController)ExecCommands(){
	id := StrParseInt(c.Input().Get("id"))
	commandstr := c.Input().Get("commands")
	cmdlist := strings.Split(commandstr,";")
	sshHost := models.GetSshHost(id)
	session ,_:= models.Connect(sshHost)
	ret := models.ExecCommands(sshHost, session, cmdlist)
	fmt.Println(ret)
	models.AddSshResult(ret)
	url := "/v1/linux/ops?id=" + c.Input().Get("id")
	
	c.Redirect(url,302)
}
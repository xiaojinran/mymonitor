package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/utils"
	"strconv"
)

func GetToken() (*utils.QyToken){
	i ,_ := strconv.ParseInt(beego.AppConfig.String("agentid"), 10, 64)
	qyinfo := &utils.QyInfo{
		Corpid:beego.AppConfig.String("corpid"),
		Corpsecret:beego.AppConfig.String("corpsecret"),
		Agentid: int(i),
	}
	token := utils.GetQyToken(qyinfo)
	return token
}

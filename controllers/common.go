package controllers

import (
	"github.com/astaxie/beego"
	"mymonitor/models"
	"strconv"
)

//指示是否弹出警告框
var alert bool
//弹出警告框内容
var alertmsg string

//获取企业微信通知令牌
func GetToken() (*models.QyToken){
	i ,_ := strconv.ParseInt(beego.AppConfig.String("agentid"), 10, 64)
	qyinfo := &models.QyInfo{
		Corpid:beego.AppConfig.String("corpid"),
		Corpsecret:beego.AppConfig.String("corpsecret"),
		Agentid: int(i),
	}
	token := models.GetQyToken(qyinfo)
	return token
}

//字符串转数字
func StrParseInt(s string) int{
	i ,_ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func StrParseFloat(s string) float32{
	i ,_ := strconv.ParseFloat(s,32)
	return float32(i)
}

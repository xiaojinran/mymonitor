package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id int `json:"id"`
	User string `json:"user"`
    Msg string `json:"msg" orm:"size(500)"`
	Time time.Time `json:"time"`
}
func AddMessage(message *Message){
	ormer := orm.NewOrm()
	ormer.Insert(message)
}

func DelMessage(message *Message){
	ormer:=orm.NewOrm()
	ormer.Delete(message)
}

func GetAllMessageByUsername(username string) ( []*Message){
	var messages  []*Message
	ormer := orm.NewOrm()
	ormer.QueryTable("message").Filter("User", username).All(&messages)
	return messages
	
}
func GetAllMessage() ( []*Message){
	var messages  []*Message
	ormer := orm.NewOrm()
	ormer.QueryTable("message").All(&messages)
	return messages
	
}


func LogMessage (message *Message){
	AddMessage(message)
}
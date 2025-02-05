package models

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type QyInfo struct {
	Corpid     string `json:"corpid"`
	Corpsecret string `json:"corpsecret"`
	Agentid int `json:"agentid"`
}


//企业微信的令牌返回结构体
type QyToken struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Expries_in int `json:"expires_in"`
	Agentid int `json:"agentid"`
}

type Content struct {
	Content string `json:"content"`
}

//微信提交内容结构体
type AppMsg struct {
	ToUser  string `json:"touser,omitempty"`
	ToParty string `json:"toparty,omitempty"`
	ToTag   string `json:"totag,omitempty"`
	Toall   string `json:"toall,omitempty"`
	MsgType string `json:"msgtype,omitempty"`
	Agentid int    `json:"agentid,omitempty"`
	Text *Content `json:"text,omitempty"`
	Safe    int    `json:"safe,omitempty"`
}


type QywechatUser struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
}

type QywechatUserlist struct {
	UserList []*QywechatUser `json:"userlist"`
}

type QywebchatRet struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	InvalidUser string `json:"invaliduser"`
}


func GetQyToken(info *QyInfo)(token *QyToken){
	url:= fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",info.Corpid,info.Corpsecret)
	httpClient:=&http.Client{}
	resp, _ := httpClient.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	json.Unmarshal(bytes,&token)
	token.Agentid = info.Agentid
	
	return
}

func SendToUser(user string, msg string, token *QyToken) (ret *QywebchatRet){
	content := &Content{Content:msg}
	appmsg := &AppMsg{
		ToUser:  user,
		MsgType: "text",
		Agentid: token.Agentid,
		Text: content,
	}
	ret = SendData(appmsg,token)
	return
	
}

func GetAllWXUser(token *QyToken) (Qy *QywechatUserlist){
	url := "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token="+ token.Access_token + "&department_id=1"
	httpclient := &http.Client{}
	resp, _ := httpclient.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&Qy)
	return
}

func SendData(msg *AppMsg, token *QyToken) (ret *QywebchatRet){
	httpclient := &http.Client{}
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token.Access_token
	bytes, _ := json.Marshal(msg)
	fmt.Println(string(bytes))
	fmt.Println(string(url))
	resp, _ := httpclient.Post(url, "application/json", bytes2.NewBuffer(bytes))
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&ret)
	return
}


package models

import "testing"

func TestGetQyToken(t *testing.T) {
	qyinfo := &QyInfo{
		Corpid:"ww1773d8b72a5648ff",
		Corpsecret:"SK8xLJ9AnUCXYYhLucUWZqjb3Nk4Mh-xjvYd_XJcHT0",
		Agentid:1000003,
	}
	qyToken := GetQyToken(qyinfo)
	t.Log(qyToken.Access_token,qyToken.Agentid)
}


func TestSendToUser(t *testing.T) {
	qyinfo := &QyInfo{
		Corpid:"ww1773d8b72a5648ff",
		Corpsecret:"SK8xLJ9AnUCXYYhLucUWZqjb3Nk4Mh-xjvYd_XJcHT0",
		Agentid:1000003,
	}
	qyToken := GetQyToken(qyinfo)
	SendToUser("xiaojinran","这是第一条测试信息",qyToken)
	SendToUser("xiaojinran","这是第二条测试信息",qyToken)
	SendToUser("xiaojinran","这是第三条测试信息",qyToken)
}

func TestGetAllWXUser(t *testing.T) {
	qyinfo := &QyInfo{
		Corpid:"ww1773d8b72a5648ff",
		Corpsecret:"SK8xLJ9AnUCXYYhLucUWZqjb3Nk4Mh-xjvYd_XJcHT0",
		Agentid:1000003,
	}
	qyToken := GetQyToken(qyinfo)
	qylist := GetAllWXUser(qyToken)
	for k,v :=range qylist.UserList {
		t.Log(k,v)
	}
}
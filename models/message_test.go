package models

import (
	"testing"
	"time"
)

func TestLogMessage(t *testing.T) {
	
	m := &Message{
		User:"xiaojinran",
		Msg:"测试消息",
		Time:time.Now(),
	}
	LogMessage(m)
}

func TestGetAllMessageByUsername(t *testing.T) {
	var msgs  []*Message

	msgs=GetAllMessageByUsername("xiaojinran")
	if msgs == nil{
		t.Log("Errors")
	}else{

		for k,v :=range msgs {
			t.Log(k,v)
		}
	}

}

package models

import (
	"bytes"
	"strings"
	"testing"
)



func Test_SSH(t *testing.T) {
	cmd := "pwd;sleep;exit"
	sshhost := &SshHost{
		Username:"root",
		Password:"XIAOJINRAN",
		Host:"119.23.66.185",
		Port:16255,
	}

	session, err := Connect(sshhost)

	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()
	
	cmdlist := strings.Split(cmd, ";")
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}
	
	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt
	
	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cmdlist {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
		
	}
	session.Wait()
	t.Log((outbt.String() + errbt.String()))
	return
}

func Test_GetSshHost(t *testing.T){
	s := GetSshHost(5)
	t.Log(s.Username)
}
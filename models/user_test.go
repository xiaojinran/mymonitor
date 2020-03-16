package models

import (
	"fmt"
	"testing"
)

func TestGetAllUser(t *testing.T) {
    var users []*User
	users = GetAllUser()
	for _,v :=range users{
		fmt.Println(v.Name , v.Pass)
	}
	
	
}


func TestAddUser(t *testing.T) {
	user:= User{
		Name:"xiaojinran",
		Pass:"password",
	}
	AddUser(&user)
}

func TestGetUserByName(t *testing.T) {
	u:=GetUserByName("xiaojinran")
	fmt.Println(u.Name , u.Pass)
	
}
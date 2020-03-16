package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
	Pass string `orm:"size(100)"`
}

func AddUser(user *User){
	ormer := orm.NewOrm()
	ormer.Insert(user)
	
}

func GetUserByName(name string) (user *User){
	u := User{Name:name}
	ormer := orm.NewOrm()
	ormer.Read(&u,"Name")
	return &u
}
func GetAllUser() ([]*User){
	var users []*User
	o :=orm.NewOrm()
	num, err := o.QueryTable("user").All(&users)
	fmt.Println(num)
	fmt.Println(err)
	return users
}
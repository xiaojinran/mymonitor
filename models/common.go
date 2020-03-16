package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	RegisterDatabase()
}

func RegisterDatabase() {
 //注册数据库，建立表，注册模型
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "itmonitor.db")
	orm.RegisterModel(new(User),new(Message))
	orm.RunSyncdb("default",false ,true)
	
	
}

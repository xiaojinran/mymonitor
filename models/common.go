package models

import (
	"bytes"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
)

func init() {
	RegisterDatabase()
}

func RegisterDatabase() {
 //注册数据库，建立表，注册模型
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "itmonitor.db")
	orm.RegisterModel(new(User),new(Message),new(SshHost),new(SshResult),new(Product))
	orm.RunSyncdb("default",false ,true)
	
	
}

// transform GBK bytes to UTF-8 bytes
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func strToFloat(str string) (f float32){
	f64 ,_ := strconv.ParseFloat(str,32)
	f = float32(f64)
	return
}


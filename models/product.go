package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Product struct {
	Id int
	Name string
	Type string
	Platform string
	Url string
	Price float32
	Time time.Time
}

func NewProduct(n string,t string,pf string,url string,p float32) *Product{
	return &Product{
		Name:n,
		Type:t,
		Platform:pf,
		Url:url,
		Price:p,
		Time:time.Now(),
		
	}
}
func AddProduct(product *Product){
	ormer := orm.NewOrm()
	ormer.Insert(product)
}

func DeleteProduct(product *Product){
	ormer := orm.NewOrm()
	ormer.Delete(product)
}

func GetAllProduct()[]*Product{
	var products []*Product
	ormer := orm.NewOrm()
	ormer.QueryTable("product").All(&products)
	return products
}
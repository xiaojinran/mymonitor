package models

type Sales struct {
	Catalog []string `json:"catalog"`
	Count []int `json:"count"`
}

func NewSales (catalog []string,count []int) (sales *Sales){
	return &Sales{
		Catalog:catalog,
		Count:count,
	}
}
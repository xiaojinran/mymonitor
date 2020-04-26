package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Stock struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Start float32 `json:"start"`
	Current float32 `json:"current"`
	Top float32 `json:"top"`
	Down float32 `json:"down"`
	Time string `json:"time"`
}

func GetStockMsg(stockid string) (stock *Stock,err error){
	stock = new(Stock)
	httpclient := &http.Client{}
	url := "http://hq.sinajs.cn/list=" + stockid
	resp, _ := httpclient.Get(url)
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	b, _ := GbkToUtf8(body)
	split := bytes.Split(b, []byte(","))
	stock.Id = stockid
	stock.Name = string(bytes.Split(split[0],[]byte("\""))[1])
	if stock.Name == ""{
		return nil, errors.New("No such a stock!")
	}
	stock.Start = strToFloat(string(split[1]))
	stock.Current = strToFloat(string(split[2]))
	stock.Top = strToFloat(string(split[3]))
	stock.Down = strToFloat(string(split[4]))
	stock.Time = string(split[30]) + " " + string(split[31])
	return stock,nil
	
}

func ConverToStock(list []string)(stock *Stock,err error){
	stock,err = mapToStock(listToMap(list))
    return

}

func mapToStock(m map[string]interface{}) (stock *Stock,err error){

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil,err
	}
	json.Unmarshal(bytes,&stock)
	return stock,nil
}

func listToMap(l []string) (map[string]interface{}){
	m := make(map[string]interface{})
    m["id"] = l[0]
    m["name"] = l[1]
    m["start"] = l[2]
    return m
}
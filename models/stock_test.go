package models

import (
	"fmt"
	"testing"
)

func TestMapToStock(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "test"
	m["top"] = 12.3
	s,err := mapToStock(m)
	if err != nil {
		t.Fail()
	}
	t.Log(s.Name,s.Top)
}

func TestListToMap(t *testing.T){
	l := []string{"sz00002","万科A","23.33"}
	m := listToMap(l)
	fmt.Println(m["name"])
}

func TestConverToStock(t *testing.T) {
	l := []string{"sz00002","万科A","23.33"}
	s ,_:= ConverToStock(l)
	t.Log(s.Name)
}


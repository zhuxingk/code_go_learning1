package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"_"`          //在使用JSON编码时，此编码不参与
	Age     int    `json:"age,string"` //将age 变为string类型
	Subject string `json:"subject_name"`
	gender  string
}

func main() {
	t1 := Teacher{
		Name:    "mary",
		Age:     25,
		Subject: "math",
		gender:  "female",
	}
	fmt.Println("t1", t1)
	encodeinfo, err := json.Marshal(&t1)
	if err != nil {
		fmt.Println("encode err:", err)
		return
	}
	fmt.Println("encodeinfo:", string(encodeinfo))

}

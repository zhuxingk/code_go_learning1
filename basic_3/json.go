package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id     int
	Name   string
	Age    int
	gender string
}

func main() {
	bob := student{
		Id:     01,
		Name:   "bob",
		Age:    20,
		gender: "male",//小写编码会在json中不会识别，必须用大写编码方式
	}

	//发送数据
	//编码，序列化，结构体——>字符串
	encode, err := json.Marshal(&bob)
	if err != nil {
		fmt.Println("encode err:", err)
		return
	}
	fmt.Println("encodeinfo", string(encode))
	//接收数据
	//反编码，反序列化
	//字符串--》结构体
	var Bob student
	if err := json.Unmarshal([]byte(encode), &Bob); err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("name:", Bob.Name)
	fmt.Println("age:", Bob.Age)
	fmt.Println("Id", Bob.Id)
	fmt.Println("gender", Bob.gender)

}

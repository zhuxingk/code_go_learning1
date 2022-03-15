package main

import "fmt"//format

func main() { //ctrl +alt +l 快速格式化代码
	var name string
	name = "pivot"
	var age int
	age = 20
	fmt.Println(name, age)//格式化输出

	var gender = "man" //直接赋值
	fmt.Println("gender", gender)

	address := "qinghai" //定义赋值，goland进行推导
	fmt.Println("address", address)
	test("grade", 50) //灰色部分表示形参
}
func test(a string, b int) {
	fmt.Println(a)
	fmt.Println(b)
}

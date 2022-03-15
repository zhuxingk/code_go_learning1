package main

import "fmt"

func main() {
	p1 := testptr_escape()
	fmt.Println("ptr", *p1)
}

//定义函数，返回string类型的指针，返回写在参数列表后面
func testptr_escape() *string {
	name := "bob"
	p0 := &name
	fmt.Println("*p0",*p0)
	city := "xining"
	ptr := &city
	return ptr
}
//use `go build -0 test.exe -gcflags "-m -m -l" escape memeory.go >1.txt 2>&1`observe memory escape,使用重定向将输出到文本文件中
//grep -E "name|city" 1.txt --color

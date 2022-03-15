package main

import "fmt"

func main() {
	name := "bob"
	ptr := &name                     //go语言使用指针时，不需要释放内存，会使用GC（garbage collector）机制；C语言不允许返回栈上的指针但是go语言可以
	fmt.Println("ptr's value", *ptr) //程序会在编译的时候确认分配的位置，编译的时候有必要会优先分配到堆上
	fmt.Println(" real ptr", ptr)

	name1 := "lily"
	ptr1 := &name1
	fmt.Println("name_1", *ptr1)
	fmt.Println("name_1", ptr1)

	//使用new关键字定义
	ptr2name := new(string)
	*ptr2name = "bob"

	fmt.Println("ptr2name", *ptr2name)
	fmt.Println("name2name", ptr2name)
	//可以返回栈上的指针，
	res := testptr()
	fmt.Println("city", *res, "address", res)
 if res == nil{
	 fmt.Println("res is nil")//nil is empty
 }else{
	 fmt.Println("res is not nil")
 }
}

//定义函数，返回string类型的指针，返回写在参数列表后面
func testptr() *string {
	city := "xining"
	ptr := &city

	return ptr

}

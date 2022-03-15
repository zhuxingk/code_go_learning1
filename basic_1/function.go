package main

import "fmt"

//函数返回值在参数列表之后
//多个返回值，需要用（）包裹，使用，分隔;单个返回值时，不用使用（）
func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}
func test2(a, b int, c string) (res int, str string, bl bool) {
	//返回值可以进行命名，可以在后续的逻辑中使用返回值的变量名进行逻辑运算
	res = a + b
	str = "hello"
	bl = true
	return //res,str, true//返回值命令后，可以直接return即可，不用同上加上变量名
}
func test3(a int) int {
	return a
}
func main() {
	a1, b1, s1 := test1(4, 9, "hello")
	fmt.Println("a1:", a1, "b1:", b1, "s1:", s1)

	a2, b2, s2 := test2(100, 1000, "eorld")
	fmt.Println("a2:", a2, "b2:", b2, "s2:", s2)
	a3 := test3(10)
	fmt.Println("a3", a3)
}

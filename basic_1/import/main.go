//package _import
package main

//同包名下的函数可以直接调用，同层级目录下不允许出现多个包名
import (
	"fmt"
	"awesomeProject/basic_1/import/add"
	"awesomeProject/basic_1/import/sub"
)

func main() {
	res1 := sub.Sub(20, 50) //包名.函数去调用
	res2 := add.Add(90, 90) //调用的函数名首字母必须大写，对外提供的函数包名，内部子函数可以小写，小写只有同包名的才可以使用
	fmt.Println("sub", res1, "add", res2)
}

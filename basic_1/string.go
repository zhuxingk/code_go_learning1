package main

import "fmt"

func main(){
	//定义
	name := "bob"
	//需要换行，输出原生字符串时，使用``;
	usage := `./a.out <option>
               -h help
               -a all`

	fmt.Println("name",name)
	fmt.Println("usage",usage)

	//长度，访问
	//string没有length 方法，使用自有函数len()进行处理
	l1 := len(name)//https://humanbenchmark.com/tests/reactiontimehttps://humanbenchmark.com/tests/reactiontime
	fmt.Println("l1",l1)
	//不需要（）
	for i := 0; i<len(name);i++{
		fmt.Printf("i:%d,v:%c\n",i,name[i])
	}
	//字符串拼接
	i,j,k:="hello ", "my" ," world"
	fmt.Println("i+j+k",i+j+k)

	//使用const修饰为常量，不能修改
	//编译前就确定了类型，预处理
	const address = "beijing"
	const test = "100"
	fmt.Println("address",address)
	fmt.Println("test",test)
}

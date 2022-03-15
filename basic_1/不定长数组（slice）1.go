package main

import (
	"fmt"
)

func main() {
	names := []string{"beijing", "shanghai", "guangzhou", "shengzheng", "hainan", "shaanxi", "Tibet"}
	//1.基于names进行赋值
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]
	//2使用数组灵活创建
	names2 := names[0:3] //左闭右开，取左不取右
	fmt.Println("names2", names2)
	//ptr := &names2
	//fmt.Println("point",*ptr)
	//fmt.Println("ponitd",ptr)
	//3浅拷贝
	names2[2] = "world"
	fmt.Println(names2)
	fmt.Println(names)
	//4如果想让slice独立于数组，可以使用copy（）进行深拷贝
	namesCopy := make([]string, len(names))
	//names是一个数组，copy函数接受的参数类型是slice，需要使用[:]将数组变成slice
	copy(namesCopy, names[:])
	namesCopy[0] = "hongkong"
	fmt.Println("names",names)
	fmt.Println("copy",namesCopy)

	//5如果从最左边开始截取，可以使用：
	names3 := names[:5]
	fmt.Println(names3)
	//6如果从最右边开始截取
	names4 := names[2:]
	fmt.Println(names4)
	//7截取全部
	names5 := names[:]
	fmt.Println(names5)
	//8截取子串，基于字符串截取子串
	sub1 := "helloworld"[6:7]
	fmt.Println(sub1)

	//9创建空切片，明确指定slice 容量,减少内存开销，提高运行效率
	str2 := make([]string, 10, 20) //type,length,cap(option)
	fmt.Println(str2)
	fmt.Println(len(str2), cap(str2))
	str2[0] = "hello"
	str2[1] = "world"

}

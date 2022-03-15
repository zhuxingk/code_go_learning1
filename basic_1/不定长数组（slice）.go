package main

import (
	"fmt"
	"time"
)

func main() {
	//切片,slice,底层也是数组，可以支持动态改变长度
	names1 := [10]string{"beijing", "shanghai", "guangzhou", "shengzhen"} //定长
	names2 := []string{"beijing", "shanghai", "guangzhou", "shengzheng"}  //不定长
	for i, v := range names1 {
		fmt.Println("i", i, "v", v)
	}

	fmt.Println("length", len(names2), "cap", cap(names2)) //打印追加数据前的数组长度和容量

	//追加数据
	names2 = append(names2, "hainan")
	fmt.Println(names2)
	//没有新的变量的时候，需要使用赋值“=”。而不能使用“：=”
	fmt.Println("length", len(names2), "cap", cap(names2)) //打印追加数据后的数组长度和容量

	names := append(names2, "shaanxi", "Tibet")
	fmt.Println(names)
	fmt.Println("length", len(names), "cap", cap(names)) //打印追加数据后的数组长度和容量
	//对于一个切片，不仅有len(),而且还有容量的限制
	nums := []int{}
	for i := 0; i <= 100; i++ {
		nums = append(nums, i)
		time.Sleep(1000)
		fmt.Println("length", len(nums), "cap", cap(nums))
	}
}

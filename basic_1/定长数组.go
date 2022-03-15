package main

import "fmt"

func main(){
	//定义数组
	nums := [10]int{1, 2, 3, 4}//常用方式
	//遍历数组
	for i:=0;i<len(nums);i++ {
		fmt.Println("i",i,"j",nums[i])
	}
	//使用for range进行遍历
	for key, value := range(nums){//修改value的值不会对nums中的数值进行修改，value的值是通过nums赋值过来的
		value += 1
		fmt.Println("key",key,"value",value,"nums",nums[key])
	}
	//go语言中，可以使用_忽略值，可以进行使用
	for _,value := range(nums){
		fmt.Println("key","value",value)
	}

	//使用make创建数组
	
}

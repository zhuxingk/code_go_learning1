package main

import "fmt"

func main() {
	//sync.RWMutex{}
	//当涉及到多goroutine时，go语言使用管道进行处理
	//使用管道不需要进行加解锁
	//a向管道中写东西，B从管道中读数据，go会自动处理
	//创建管道：创建装数字的管道
	numChan := make(chan int)
	go func() {
		for i := 0; i <= 50; i++ {
			data := <-numChan
			fmt.Println("read data", data)
		}
		//str := <-strChan
		//fmt.Println("read string data", str)
	}()

	//strChan := make(chan string)
	//使用管道时必须make ,同map,否则管道为空
	for i := 0; i <= 50; i++ {
		numChan <- i
		fmt.Println("write data:i", i)
	}
	//strChan <- "my world"
	//fmt.Println("write string data", strChan)

}

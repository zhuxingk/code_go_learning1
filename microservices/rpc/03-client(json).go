package main

import (
	"fmt"
	"net/rpc"

	//"net/rpc"
)

func main(){
	//1.使用rpc链接服务器---net.Dial()
	//conn ,err := rpc.Dial("tcp","127.0.0.1:8089")
	//使用JSON进行编码避免出现其他地方无法解码的情况
	conn ,err :=rpc.Dial("tcp","127.0.0.1:8089")
	if err!=nil{
		fmt.Println("dail() err:",err)
		return
	}
	defer conn.Close()
	//2.调用远程函数
	var reply string//接受返回值，传出参数
	err = conn.Call("hello.Helloworld","李白",&reply)
	if err!=nil{
		fmt.Println("call err():",err)
		return
	}
	fmt.Println("reply",reply)
}


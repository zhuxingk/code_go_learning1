package main

import (
	"fmt"
)

/*func main01() {
	//1.使用rpc链接服务器---net.Dial()
	conn, err := rpc.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("dail() err:", err)
		return
	}
	defer conn.Close()
	//2.调用远程函数
	var reply string //接受返回值，传出参数
	err = conn.Call("hello.Helloworld", "李白", &reply)
	if err != nil {
		fmt.Println("call err():", err)
		return
	}
	fmt.Println("reply", reply)
}*/

func main() {
	Myclient := Initclient("127.0.0.1:8087")
	var resp string
	err := Myclient.Helloworld("杜甫", &resp)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	fmt.Println(resp, err)
}

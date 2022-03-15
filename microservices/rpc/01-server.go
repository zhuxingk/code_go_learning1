package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//定义类对象
type World struct {
}

//绑定类方法
func (this *World) Helloworld(name string, resp *string) error {
	*resp = name + "你好"
	return nil
}

func main() {
	//	1.注册RPC服务，绑定对象方法
	/*err := rpc.RegisterName("hello",new(World))
	if err != nil{
		fmt.Println("rpc服务注册失败！",err)
		return
	}*/
	Registerservice(new(World))
	//	2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8087")
	if err != nil {
		fmt.Println("listen ():err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("start to listen!")
	//	3.建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept ():err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("connection is ok")
	//	4.绑定服务
	rpc.ServeConn(conn)
}

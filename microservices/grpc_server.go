package main

import (
	"awesomeProject/microservices/PB/PB" //import package需要注意文件路径
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//定义类
type Children struct {
}

//按照接口绑定类方法
func (this *Children) SayHello(ctx context.Context, teacher *PB.Teacher) (*PB.Teacher, error) {
	teacher.Name += "is sleeping"
	//if error !=nil{
	//	fmt.Println("")
	return teacher, nil
}

func main() {
	//初始化grpc对象
	grpcServer := grpc.NewServer()
	//注册服务
	PB.RegisterSayNameServer(grpcServer, new(Children))
	//设置监听，指定IP，port
	listener, err := net.Listen("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("listen err :", err)
		return
	}
	defer listener.Close()
	//启动服务
	grpcServer.Serve(listener)
}

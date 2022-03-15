package main

import (
	"awesomeProject/microservices/PB/PB"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main()  {

	//链接grpc服务
	grpcConn,err := grpc.Dial("127.0.0.1:8089",grpc.WithInsecure())
	if err!= nil{
		fmt.Println("dail err :",err)
		return
	}
	defer grpcConn.Close()
	//初始化grpc客户端
	grpcclient := PB.NewSayNameClient(grpcConn)

	//创建并初始化teacher对象
	var teacher PB.Teacher
	teacher.Name="itcast"
	teacher.Age=20
	//远程调用服务
	t, err:=grpcclient.SayHello(context.TODO(),&teacher)
	if err!=nil{
		fmt.Println("client err:",err)
		return
	}
	fmt.Println(t)
}

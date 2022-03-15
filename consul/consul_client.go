package main

import (
	"awesomeProject/consul/PB1/PB1"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//初始化consul配置，客户端服务端需要一致
	consulconfig := api.DefaultConfig()
	//创建consul对象(可以重新指定consul的属性：ip,port etc.)
	consulclient, err := api.NewClient(consulconfig)
	if err != nil {
		fmt.Println("consul err:", err)
		return
	}
	//获取consul操作对象,获取健康服务
	services, _,err := consulclient.Health().Service("consul_test","grpc",true,nil)
    //services[0].Service.port
	if err != nil{
		fmt.Println("services err:",err)
		return
	}
	target := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	//获取地址

	//和GRPC服务建立链接

	/////////////////////////////////////////////////////////////////

	//链接服务
	//grpconn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	//使用consul上的IP、port参与服务建立链接
	grpcconn, err := grpc.Dial(target,grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	//初始化客户端
	grpcclient := PB1.NewHelloClient(grpcconn)
	var person PB1.Person
	person.Age = 18
	person.Name = "baby"
	//调用远程函数
	p, err := grpcclient.SayHello(context.TODO(), &person)
	fmt.Println(p, err)

}

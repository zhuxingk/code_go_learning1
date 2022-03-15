package main

import (
	"awesomeProject/consul/PB1/PB1"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

//定义类对象
type Children1 struct {
}

func (this *Children1) SayHello(ctx context.Context, p *PB1.Person) (*PB1.Person, error) {
	p.Name = "hello " + p.Name
	return p, nil
}
func main() {
	//grpc服务注册到consul上
	//初始化consul服务配置
	consulconfig :=api.DefaultConfig()
	//创建consul对象
	consulclient ,err := api.NewClient(consulconfig)
	if err != nil{
		fmt.Println("create err:",err)
		return
	}
	//告诉consul即将注册的信息
	req := api.AgentServiceRegistration{
		ID:      "test",
		Name:    "consul_test",
		Tags:    []string{"grpc","consul"},
		Port:    8800,
		Address: "127.0.0.1",
		Check:   &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			Name:     "test",
			Interval: "5s",
			Timeout:  "1s",
			TCP:      "127.0.0.1:8800",
		},
	}
	//注册grpc服务到consul上
	consulclient.Agent().ServiceRegister(&req)
	////////////////////////////////////////
	//初始化grpc对象
	grpcserver := grpc.NewServer()
	//注册服务
	PB1.RegisterHelloServer(grpcserver, new(Children1))
	//设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	fmt.Println("server start!")
	//启动服务
	grpcserver.Serve(listener)

}

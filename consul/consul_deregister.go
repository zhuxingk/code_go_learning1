package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
) 

func main(){
	//初始化consul配置
	consulconfig := api.DefaultConfig()
	//创建consul对象
	consulclient , err := api.NewClient(consulconfig)
	if err !=nil{
		fmt.Println("create err:",err)
		return
	}
	consulclient.Agent().ServiceDeregister("test")
}

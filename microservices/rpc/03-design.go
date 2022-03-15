package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//服务端在注册rpc对象时，能在编译期间监测到注册对象是否合法
//创建接口，在接口中定义方法原型
type Myinterface interface {
	Helloworld(string, *string) error
}

//调用该方法时，需要给i传参，参数应该是实现了helloworld的方法的类对象
func Registerservice(i Myinterface) {
	rpc.RegisterName("hello", i)
}

//---------------use for client
//像调用本地函数一样，调用远程函数
//定义类
type Myclient struct {
	c *rpc.Client
}

//由于使用了c调用call,需要对c进行初始化
func Initclient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{c: conn}
}

//实现函数，参照上面的	interface进行实现
func (this *Myclient) Helloworld(a string, b *string) error {
	//参数1参照上面的interface中的registername,a传入参数，b，传出参数
	return this.c.Call("hello.Hellowrold", a, b)

}

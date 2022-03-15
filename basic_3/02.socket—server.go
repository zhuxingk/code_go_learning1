package main

import (
	"fmt"
	"net"
	"strings"
)

func main(){
	//create listen
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d",ip,port)//attention the format
	//func Listen(network string, address string) (Listener, error)
	//net.Listen("tcp",8848)
	 listener, err :=net.Listen("tcp",address)

	 if err != nil{
		 fmt.Println("net.listen:err:",err)
		 return
	 }
	 fmt.Println("listening")
	 conn, err := listener.Accept()
	 if err != nil{
		 fmt.Println("listener.accept err:",err)
		 return
	 }
	 fmt.Println("connection is build")
	 //create container
	 buf := make([]byte,1024)//use make to create slice
	 //cnt is real data length from client
	 cnt,err :=conn.Read(buf)
	 if err!=nil{
		 fmt.Println("conn.read.err:",err)
		 return
	 }
	fmt.Println("client send to server,length:",cnt,"send to server,context",string(buf))
	 //将数据转换为大写,模拟server响应
	UPperdata := strings.ToUpper(string(buf))
	cnt,err = conn.Write([]byte (UPperdata))
	fmt.Println("server send to client,length:",cnt,"send to client,context",UPperdata[0:cnt])
	conn.Close()
}

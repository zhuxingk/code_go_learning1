package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//create listen
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port) //attention the format
	//func Listen(network string, address string) (Listener, error)
	//net.Listen("tcp",8848)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.listen:err:", err)
		return
	}
	fmt.Println("listening")

	//server can accept many connections,main goroutine ->listening;son goroutine -> data processing
	//every connection can accept and process many data requests
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept err:", err)
			return
		}
		fmt.Println("connection is build")
		//create container
		go handle(conn)
	}
}

//处理具体业务的逻辑，需要将conn传递进来，每一个新链接，conn是不同的
func handle(conn net.Conn) {
	for {
		buf := make([]byte, 1024) //use make to create slice
		//cnt is real data length from client
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.read.err:", err)
			return
		}
		fmt.Println("client send to server,length:", cnt, "send to server,context", string(buf[0:cnt]))
		//将数据转换为大写,模拟server响应
		UPperdata := strings.ToUpper(string(buf[0:cnt]))
		cnt, err = conn.Write([]byte(UPperdata))
		fmt.Println("server send to client,length:", cnt, "send to client,context", UPperdata)
		//_ = conn.Close() //ignore the error
	}
}

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")//attention the format
	if err != nil {
		fmt.Println("net.dail.err", err)
		return
	}
	fmt.Println("server and client is connected!")
	//send data
	senddata := []byte("helloworld")
	for {
		cnt, err := conn.Write(senddata)
		if err != nil {
			fmt.Println("senddata err:", err)
			return
		}
		fmt.Println("send to server:", cnt, "data:", string(senddata))

		//receive data from server
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("receive data err:", err)
			return
		}
		fmt.Println("server send to client:", cnt, "data:", string(buf[0:cnt]))

		//_ = conn.Close()
	}
}

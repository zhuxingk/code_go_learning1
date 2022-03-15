package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	//http://127.0.0.1:8088/user,func 是回调函数，用于路由的响应
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request contains data from client
		fmt.Println("user detail:")
		fmt.Println("request:", request)
		//via writer return data to client
		_, _ = io.WriteString(writer, "this is /user")

	})
	//http://127.0.0.1:8088/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "this is /name")

	})
	//http://127.0.0.1:8088/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "this is /id")

	})
	//func ListenAndServe(addr string, handler Handler) error {
	//	server := &Server{Addr: addr, Handler: handler}
	//	return server.ListenAndServe()
	//}
	if err := http.ListenAndServe("127.0.0.1:8088", nil); err != nil {
		fmt.Println("listen err:", err)
		return
	}
	//if err != nil{
	//fmt.println("listen err",err)
	//return
	//}
}

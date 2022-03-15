package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http package
	client := http.Client{}
	resp, err := client.Get("https://baidu.com")
	if err != nil {
		fmt.Println("client get err", err)
		return
	}
	body := resp.Body
	//Body io.ReadCloser
	readbodystr, err := ioutil.ReadAll(body)
	fmt.Println("body1",body)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}

	fmt.Println("body", string(readbodystr))

	ct := resp.Header.Get("content-type")
	date := resp.Header.Get("date")
	server := resp.Header.Get("server")

	fmt.Println("content type:", ct)
	fmt.Println("date", date)
	fmt.Println("server", server)
	//beego ,gin web 框架//TODO
	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	fmt.Println("url:", url)
	fmt.Println("code", code)
	fmt.Println("status", status)

}

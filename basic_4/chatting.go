package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type user struct {
	name string
	id   string
	msg  chan string
}

//创建全局map ，用于保存user数据
var allusers = make(map[string]user)

//定义message全局通道，用于接受任何用户的消息
var Message = make(chan string)
//上锁
//var lock sync.RWMutex
//将所有代码全部写在同一个文件
func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	//启动全局唯一的goroutine,负责监听message通道，写给所有用户
	go broadcast()
	fmt.Println("server start,listening...")
	for {
		fmt.Println("main goroutine is listening...")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			return
		}
		fmt.Println("server is connected!")
		//建立链接
		//启动处理业务的goroutine
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	fmt.Println("启动业务！")
	//客户端域服务器建立链接的时候，会有IP和port，用这个作为user的ID
	clientaddr := conn.RemoteAddr().String()
	//创建user
	newuser := user{
		name: clientaddr,            //初始值，后续可以修改
		id:   clientaddr,            //唯一值
		msg:  make(chan string, 20), //make空间，否则无法写入数据
	}
	//添加user到map结构
	allusers[newuser.id] = newuser
	//启动gouroutine,返回给客户端
	go writebacktoclient(&newuser, conn)
	//定义退出信号，监听client退出
	var isQuit = make(chan bool)
	//用于充值计数器，告知watch函数用户在线
	var restTimer = make(chan bool)
	//退出
	go watch(&newuser, conn, isQuit, restTimer)
	//向message写入数据，当前用户的上线消息，用于通知所有用户（广播）
	logininfo := fmt.Sprintf("[%s]:[%s]login!\n", newuser.id, newuser.name)
   // lock.Lock()
	Message <- logininfo
	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("start to exit!")
			//map删除用户，conn close
			//服务器主动退出，不进行真正的退出工作，而是发送一个退出信号，统一做退出处理，可以使用新的管道做信号传递
			isQuit <- true
		}
		if err != nil {
			fmt.Println("conn.Read.err:", err)
			return
		}
		fmt.Println("服务器接受到客户端发送的数据：", string(buf[0:cnt-1]), "cnt", cnt)
		//业务逻辑处理
		//查询当前所有用户
		//先判断接受的数据是否是who，length&string
		userInput := string(buf[0 : cnt-1])
		if len(userInput) == 3 && userInput == "who" {
			fmt.Println("查询用户信息")
			//这个切片包含所有的用户信息
			var userinfos []string
			for _, user := range allusers {
				userinfo := fmt.Sprintf("userid:%s,username:%s", user.id, user.name)
				userinfos = append(userinfos, userinfo)
			}
			//最终到管道中的是一个字符串
			r := strings.Join(userinfos, "\n")
			//将数据返回给查询的客户端
			newuser.msg <- r

			//规则：rename|bob
			//1.读取长度为7且判断字符串是rename
			//		2.使用|进行分隔，|后面的部分作为名字
			//		3.更新用户名字newuser.name=bob
			//		4.通知客户端，更新成功
		} else if len(userInput) > 8 && userInput[:6] == "rename" {
			fmt.Println("user fix name!")
			arr := strings.Split(userInput, "|")
			newuser.name = arr[1]
			allusers[newuser.id] = newuser
			//newuser.name = strings.Split(userInput, "\n")[1]
			newuser.msg <- "rename successfully!"
		} else {
			//如果用户输入的不是“who"命令，只是普通的聊天的信息，写入广播，由其他的goroutine进行常规转发
			Message <- userInput
		}
		//遍历alluser这个map，（key:userid,value:user）,拼接ID和name,返回到客户端

		restTimer <- true
		//lock.Unlock()
	}
}

//向所有用户广播消息,全局唯一
func broadcast() {
	fmt.Println("start !")
	defer fmt.Println("main goroutine is exit!")
	//读取message数据
	for {
//lock.Lock()
		info := <-Message
		fmt.Println("accept info:", info)
		//将数据写入到每一个用户的msg管道中
		for _, user := range allusers {
			user.msg <- info
			//lock.Unlock()
		}
	}
}

//每个用户应该i有一个用来监听的自己mag管道的goroutine,负责将数据返回给客户端
func writebacktoclient(user *user, conn net.Conn) {
	fmt.Println("正在监听自己的msg通道")
	for data := range user.msg {
		fmt.Printf("user的%s正在监听自己的msg channel\n", user.name)
		_, _ = conn.Write([]byte(data))
		//backdata, err := conn.Write([]byte(data))
		//if err != nil {
		//	fmt.Println("back err", err)
		//	return
		//}
		fmt.Println("backdata", data)
	}
}

//启动一个goroutine,负责监听退出信号，触发后进行清零工作，：delete map,close conn
func watch(user *user, conn net.Conn, isQuit, restTimer <-chan bool) {
	fmt.Println("start to listen exit signal")
	//lock.Lock()
	for {
		select {
		case <-isQuit:
			Logoutinfo := fmt.Sprintf("%s is exit!", user.name)
			fmt.Println("start to delete user!", user.name)
			delete(allusers, user.id)
			Message <- Logoutinfo
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			Logoutinfo := fmt.Sprintf("timeout %s is exit!", user.name)
			fmt.Println("start to delete user!", user.name)
			delete(allusers, user.id)
			Message <- Logoutinfo
			conn.Close()
			return
		case <-restTimer:
			fmt.Println("the user keep alive", user.name)

		}
		//lock.Unlock()
	}

}
//map竞争（读写）需要上锁（解锁）//TODO
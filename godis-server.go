package main

import (
	"net"
	"log"
)

//服务端，用于存储客户端发送过来的数据
func main() {
	//以tcp的方式，监听本机9736端口
	netListen, err := net.Listen("tcp","127.0.0.1:9736")
	if err != nil {
		log.Print("listen err")
	}
	//关闭客户端连接
	defer netListen.Close()

	//循环监听客戶端，发送的请求
	for{
		conn,err := netListen.Accept()
		if err != nil {
			continue
		}
		//开启一个go程，处理客户的连接请求
		go handle(conn)
	}
}

func handle(conn net.Conn){
	for{
		responseConn(conn, "accept 127.0.0.1:9736")
	}
}

func responseConn(conn net.Conn,result string) {
	conn.Write([]byte(result))
}
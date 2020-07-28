package main

import (
	"github.com/karlwanghn/gostudy/grpc"
	"log"
	"net"
	"net/rpc"
)

func main() {
	//注册远程服务
	rpc.RegisterName("HelloService", new(HelloService))
	//设置监听端口是
	listener, err := net.Listen("tcp", ":1234")
	//如果返回错误，则进行处理
	if err != nil {
		log.Fatal("无法正常访问")
	}
	//接收请求
	conn, eer := listener.Accept()
	if eer != nil {
		log.Fatal("接收请求出错")
	}
	rpc.ServeConn(conn)
	log.Println("服务请求成功")
}

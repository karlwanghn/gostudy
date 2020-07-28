package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//定义一个服务结构体
type HelloService struct {
}

/**
*大写的Hello表式 必须要大写，因为此方法要公开，不能小写
*定义远程方法Hello
 */
func (p *HelloService) Hello(request string, response *string) error {
	fmt.Println("******服务端被调用************")
	*response = "hello:" + request
	return nil
}

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
	fmt.Println("是否在方法调用之前")
	rpc.ServeConn(conn)
	log.Println("服务请求成功")

}

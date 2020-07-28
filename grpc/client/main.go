package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("连接异常")
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello", reply)
	if err != nil {
		log.Fatal("调用异常")
	}
	fmt.Print("打印请求返回：" + reply)
}

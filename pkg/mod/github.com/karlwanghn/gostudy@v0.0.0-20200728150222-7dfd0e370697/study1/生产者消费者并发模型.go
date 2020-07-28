package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者: 生成 factor 整数倍的序列 2.
func Producer(factor int, out chan<- int) {
	fmt.Println("生产者进行生产")
	for i := 0; ; i++ {
		out <- i * factor
		fmt.Println("生产的结果%d", i*factor)
	}
	fmt.Println("生产的结果" + string(factor))
}

// 消费者
func Consumer(in <-chan int) {
	fmt.Println("消费者进行消费")
	for v := range in {
		fmt.Println("消费者进行消费d%", v)
	}
}
func main() {
	ch := make(chan int, 64)
	fmt.Println(cap(ch))
	// 成果队列
	go Producer(3, ch)
	// 生成 3 的倍数的序列
	go Producer(5, ch)
	// 生成 5 的倍数的序列
	go Consumer(ch)
	// 消费 生成的队列
	//运行一定时间后退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)

}

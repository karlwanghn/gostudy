package main

import (
	"fmt"
)

//执行时出现异常 API server listening at: 127.0.0.1:54399
//为什么执行会出现有时可以打印，有时不可以打印
//采用互斥锁的方式是比较low的，go语言经典编程里说不要用共享内存来处理并发，要用通信来进行内存共享
/*func main() {
var mu sync.Mutex
mu.Lock()
go func(){
	fmt.Println("你好, 世界")
	mu.Unlock()
}()
mu.Lock()*/

func main() {
	//创建一个管道
	done := make(chan int)
	go func() {
		fmt.Println("你好, 世界")
		done <- 1
	}()

	fmt.Println("主函数执行完了，主线程消亡")
	<-done
}

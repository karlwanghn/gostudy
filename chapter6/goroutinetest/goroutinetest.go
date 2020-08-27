package main

import (
	"fmt"
	"runtime"
)

/**
并发的简单模型
*/
func say(hello string) {
	for i := 0; i < 5; i++ {
		//让出cup给其它gorountine
		runtime.Gosched()
		fmt.Println(hello)
	}
}

func main() {
	//runtime.NumCPU()获取当前主机的cpu数量
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	//当cpu数量为2时，二个方法可以同时执行完成，主进程才完成
	//当cpu数据量1时，只有主默认能完成，主线程就结束，等待线程无法正常完成程序结束
	//开启一个新的goroutine
	go say("world")
	//当前默认goroutine执行
	say("hello1")

}

package main

import "fmt"

func Count(ch chan int, num int) {
	fmt.Println("程序进行设置", num)
	ch <- num

}
func main() {
	//创建通道数组，长度为10
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		//为每个goroutine创建一个通道
		fmt.Println(i, "通道创建开始")
		go Count(chs[i], i)
		fmt.Println(i, "通道创建完成")
	}

	fmt.Println("程序进入从通道取值")
	//从每个通道中取值
	for _, ch := range chs {
		fmt.Println("取值程序进入阻塞，当设置后才可以取值")
		a, _ := <-ch
		fmt.Println(a, "程序进行取值完成")

	}
}

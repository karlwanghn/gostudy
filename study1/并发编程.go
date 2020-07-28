package main

var a string

func f() {
	print(a)
}

//当启动一个新的goroutine时，不能保证函数一定就是顺序执行，
//因为已经启动了一个新的goroutine,属于不同的cup去执行相应的程序
//但是这二个goroutine是属于一个主goroutine来管理，多个goroutine之间是相互通信的
func hello() {
	a = "hello, world"
	go f()
}

func main() {
	hello()
}

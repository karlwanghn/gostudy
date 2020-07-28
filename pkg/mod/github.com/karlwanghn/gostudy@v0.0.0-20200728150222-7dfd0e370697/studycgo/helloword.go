package main

/*
	#include <stdio.h>
	static void SayHello(const char* s) {
		puts(s);
   }
*/
import "C"
import (
	"fmt"
)

func main() {
	C.puts(C.CString("Hello, World 1\n"))
	fmt.Println("Hello world!2")
	C.SayHello(C.CString("Hello, World 3\n"))

}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("调用失败", err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

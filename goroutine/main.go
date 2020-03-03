package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang 主线程" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

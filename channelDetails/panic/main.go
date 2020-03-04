package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Print("主线程在执行中....\n")
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}

package main

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("WriteData写入的值为=%v\n", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		} else {
			fmt.Printf("readData读到的值为=%v \n", v)
			time.Sleep(time.Second)
		}
	}

	//读取完数据, 需要给chan写一个标识位
	exitChan <- true
	close(exitChan)

}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 11)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

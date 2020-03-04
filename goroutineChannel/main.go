package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从intChan取， 存入primeChan， 结束的时候给exitChan发一个信号
func CalPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		flag = true
		num, ok := <-intChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("当前协程取完数据...")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 8000) // 放入结果的管道
	exitChan := make(chan bool, 4)    //标识管道

	//开启协程，放入数据, 并计算
	start := time.Now().UnixNano()

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go CalPrime(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().UnixNano()
		fmt.Printf("使用协程耗时:%d", end-start)
		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d ", res)
	}

	fmt.Println("主线程退出")

}

package main

import "fmt"

func main() {
	//var chan1 chan int //双向管道 可读可写

	var chan2 chan<- int //只写管道

	chan2 = make(chan int, 5)

	chan2 <- 20
	chan2 <- 30
	chan2 <- 40
	chan2 <- 50
	//num := <-chan2 error receive from send-only type chan<- int
	fmt.Println(chan2)

	var chan3 <-chan int //只读管道
	num3 := <-chan3
	//chan3 <- 30 // error send to receive-only type <-chan int
	fmt.Println(chan3)
}

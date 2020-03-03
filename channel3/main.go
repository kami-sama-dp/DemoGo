package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	// intChan <- 300
	close(intChan)

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}
}

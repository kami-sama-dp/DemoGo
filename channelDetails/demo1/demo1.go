package main

import "fmt"

func main()  {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	strChan := make(chan string , 5)
	for i := 0; i < 5; i++ {
		strChan<- "hello" + fmt.Sprintf("%d", i)
	}
	// label:
	for{
		select{
		case v:= <-intChan:
			fmt.Printf("从intChan读取数据---%d\n", v)
		case v:= <- strChan:
			fmt.Printf("从strChan读取数据---%s\n", v)
		default:
			fmt.Println("取不到....")
			return
			// break label
		}
	
	}
}
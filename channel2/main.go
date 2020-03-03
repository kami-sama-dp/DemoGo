package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Tom jack"
	cat := Cat{
		Name: "cat_name",
		Age:  12,
	}
	allChan <- cat
	<-allChan
	<-allChan
	data := <-allChan
	fmt.Printf("%T\n", data)
	a := data.(Cat) //类型断言
	fmt.Println(a.Age)
}

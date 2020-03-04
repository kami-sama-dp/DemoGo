package main

import (
	"fmt"
)

func returnValue() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}

func main() {
	a := returnValue()
	fmt.Println(a)
}

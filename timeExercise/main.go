package main

import (
	"fmt"
	"strconv"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T, num1的值为%v,num1的地址%v\n ", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2的类型%T, num2的值为%v,num2的地址%v\n ", num2, *num2, &num2)
	// start := time.Now().Unix()
	// test03()
	// end := time.Now().Unix()
	// fmt.Printf("执行test03花费时间为 %v\n", end-start)
}

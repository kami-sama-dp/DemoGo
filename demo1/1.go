package main

import (
	"fmt"
	"strings"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n=", n)
	}
}

func AddUpper() func(int) int {
	a := 10
	return func(x int) int {
		a += x
		return a
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println("文件名", f("test"))
	fmt.Println("文件名", f("demo.jpg"))
	fmt.Println("文件名", f("demo.avi"))
}

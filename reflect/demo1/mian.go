package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("%T\n", rVal)

	iVal := rVal.Interface()
	iVal = iVal.(int)
	fmt.Println(iVal)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	//	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	iVal := rVal.Interface()
	fmt.Printf("%v %T\n", iVal, iVal)

}

func main() {
	var num int = 100

	reflectTest01(num)
	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
}

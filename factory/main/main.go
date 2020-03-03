package main

import (
	"DemoGo/factory/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("tom~", 88.8)
	fmt.Println(stu)
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.GetScore())
}

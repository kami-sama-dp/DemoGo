package main

import (
	"DemoGo/encapsulate/model"
	"fmt"
)

func main() {
	person := model.NewPerson("Tom")
	person.SetAge(12)
	person.SetSal(6000)
	fmt.Println(person)
	fmt.Println(person.Name, person.GetAge(), person.GetSal())
}

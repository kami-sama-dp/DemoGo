package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

type interger int

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster hello")
}

func (m Monster) Say() {
	fmt.Println("monster SAy()")
}

func (i interger) Say() {
	fmt.Println("222222    ", i)
}

func (stu Stu) Say() {
	fmt.Println("11111")
}

func main() {
	var monster Monster
	var a AInterface = monster
	a.Say()
	var b BInterface = monster
	b.Hello()
	fmt.Println("111", monster)
}

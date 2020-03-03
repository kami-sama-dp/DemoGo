package main

import "fmt"

type Monkey struct {
	Name string
}

type LittleMonkey struct {
	Monkey //继承
}
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树.....")
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，会飞翔了.....")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习，会游泳了....")
}

func main() {
	littleMonkey := LittleMonkey{
		Monkey{
			Name: "Tom",
		},
	}
	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}

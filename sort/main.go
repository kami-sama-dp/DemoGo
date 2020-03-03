package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) say() string {
	info := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		s.name, s.gender, s.age, s.id, s.score)
	return info
}

type Box struct {
	length float64
	width  float64
	height float64
}

func (box *Box) getVolum() float64 {
	return box.length*box.height + box.width
}

type Visitor struct {
	Age  int
	Name string
}

func (visitor *Visitor) showPrice() {
	if visitor.Age > 90 || visitor.Age < 8 {
		fmt.Println("考虑到安全， 请不要游玩")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v  年龄为%v  收费20元\n",
			visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为 %v  年龄为%v  免费\n",
			visitor.Name, visitor.Age)
	}
}

func main() {
	var stu = Student{
		name:   "Tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	str := stu.say()
	fmt.Println(str)

	var box = Box{
		width:  14.5,
		length: 12.3,
		height: 5.2,
	}
	str1 := box.getVolum()
	fmt.Printf("体积为%.2f\n", str1)

	var v Visitor
	for {
		fmt.Println("请输入你的名字:")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄:")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}

package main

import (
	"fmt"
)

//go语言中用结构体模拟实现"继承"

type animal struct {
	name string
}

func (ani *animal) move() {
	fmt.Printf("%s开始行动\n", ani.name)
}

type dog struct {
	foot uint8
	animal	//这就算是个继承，animal的
}

func (d dog) wang() {
	fmt.Printf("%s在\"汪汪\"叫\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "狗1",
		},
		foot: 4,
	}

	fmt.Println(d1)
	d1.wang()
	d1.move()
}

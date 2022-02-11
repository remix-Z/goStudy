package main

import "fmt"

//一个类型可以实现多个接口
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
}

func (c *cat) eat(food string) {
	fmt.Printf("%s吃%s\n", c.name, food)
}

func (c *cat) move() {
	fmt.Println("猫移动")
}

func order(a animal, food string) {
	a.move()
	a.eat(food)
}

func main() {
	c1 := cat{
		"小黑",
	}

	var m mover
	var e eater
	var a animal
	m = &c1
	e = &c1
	a = &c1

	fmt.Println(m)
	fmt.Println(e)
	fmt.Println(a)

	order(a, "猫粮")

}

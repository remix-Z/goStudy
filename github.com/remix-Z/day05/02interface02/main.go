package main

import "fmt"

type animal interface {
	move()
	speak()
}

type cat struct {
	name string
}

//指针接收者
func (c *cat) speak() {
	fmt.Println("喵喵喵")
}

func (c *cat) move() {
	fmt.Println("干饭")
}

func do(a animal) {
	a.move()
	a.speak()
}
func main() {
	var a1 animal
	var a2 animal
	c1 := cat{
		"小黑",
	}

	c2 := &cat{
		"小刘",
	}

	do(&c1)
	fmt.Printf("%T\n", a1) //nil
	//a1 = c1	如果使用指针接收者,变量需要取地址
	a1 = &c1
	fmt.Println(a1)
	fmt.Printf("%T\n", a1) //cat
	//接口本身保存 值类型和值本身

	a2 = c2
	fmt.Println(a2)
	fmt.Printf("%T\n", a2) //*main.cat
}

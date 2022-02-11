package main

import "fmt"

//接口 接口是一种类型

//引出接口的实例

type cat struct{}

type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵")
}

func (d dog) speak() {
	fmt.Println("汪")
}

//无论传什么类型的参数
func da(x speaker) {
	//接收一个参数 传进来什么就打什么
	x.speak()
}

//定义一个能speak的类型
//特殊的类型 只规定实现了什么方法
type speaker interface {
	speak()
}

//编程中会遇到以下场景 不关心对象类型 只关心方法
func main() {
	var c1 cat
	var d1 dog

	da(c1) //c1.speak()
	da(d1) //d1.speak()
}

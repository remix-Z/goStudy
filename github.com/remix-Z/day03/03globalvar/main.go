package main

import (
	"fmt"
)

//变量的作用域

var x = 100 //定义一个全局变量
func f1() {
	//函数中查找变量的顺序
	//1 先在函数内部查找
	//2 如果找不到就在函数外部查找
	x := 5
	//name := "小明"
	fmt.Println(x)
}

func main() {
	f1()
	//fmt.Println(name)

	//语句块作用域
	for i := 0; i < 10; i++ {

	}
	//fmt.Println(i) i的作用域再for循环里
}

package main

import "fmt"

//闭包的例子
func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

//变动
func adder02(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret01 := adder()   //adder()返回值是func(int) int类型，用ret01去接收
	ret02 := ret01(20) //ret01是func(int) int类型,ret02去接收func(int) int的返回值
	fmt.Println(ret02) //100+20 = 120

	ret01 = adder02(100)
	ret02 = ret01(20)
	fmt.Println(ret02)
}

//闭包	???
//闭包是一个函数，这个函数包含了外部函数的变量
//底层原理:
//1、函数可以作为返回值
//2、函数内部查找变量顺序，现在内部，后在外部

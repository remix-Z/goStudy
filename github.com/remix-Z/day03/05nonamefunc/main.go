package main

import (
	"fmt"
)

//匿名函数
var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)

	//函数内部不能声明带名字的函数，所以一般用在函数内部
	f2 := func(x, y int) {
		fmt.Println(x - y)
	}

	f2(10, 2)

	//如果只执行一次，可以简写成立即执行的函数
	func(x, y int) {
		fmt.Println(x, y)
		fmt.Println("hello world")
	}(100, 200)
}

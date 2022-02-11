package main

import (
	"fmt"
)

//闭包
func f1(f func()) {
	fmt.Println("this is function 01")
	f()
}

func f2(x, y int) {
	fmt.Println("this is function 02")
	fmt.Println(x, y)
}

func f3(f func(int, int), m, n int) func() {
	return func() {
		fmt.Println("hello")
		f(m, n)
	}
}

func main() {
	//闭包的应用，想把f2当做参数传到f1函数中
	//定义一个f3对f2进行包装
	ret := f3(f2, 100, 200)
	f1(ret)
}

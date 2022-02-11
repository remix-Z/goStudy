package main

import (
	"fmt"
)

//函数类型
func f1() {
	fmt.Println("hello world")
}

func f2() int {
	return 10
}

//函数类型可以当做参数
func f3(x func() int) {
	fmt.Println(x())
}

func f4(x, y int) int {
	return x + y
}

//函数作为返回值
func f5(x func() int) func(int, int) int {
	return f4
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) //func()

	b := f2
	fmt.Printf("%T\n", b)  //func() int
	fmt.Printf("%T\n", f4) //func(int,int) int

	f3(f2)
	//f3(f4(1, 2)) 	因为f4有参数所以不满足f3的参数要求

	fmt.Println(f5(f2))        //输出是地址
	fmt.Printf("%T\n", f5(f2)) //f4类型

	fxxxx := f5(f2)
	fmt.Println(fxxxx(1, 1))

	//想知道这使用的目的呢??? 递归？
}

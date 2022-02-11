package main

import (
	"fmt"
)

//函数的意义
//一段代码的封装
//一段逻辑抽象封装成一个函数，使用时直接调用函数名即可

//函数的定义
func Myfunc01(x int, y int) (ret int) {
	return x + y
}

//变种1 没有返回值的函数
func Myfunc02(x int, y int) {
	fmt.Printf("****%T****\n", x) //应该类似c，引用类型?
	fmt.Println(x + y)
}

// func Myfunc03() (ret int) {	可以简化写法 返回值可以命名也可以不命名
// 命名即声明
func Myfunc03() string {
	return "我爱北京天安门"
}

//多个返回值
func Myfunc04() (int, int) {
	return 1, 2
}

//参数类型简写
//连续多个参数类型相同时，除最后一个的可以不写
func Myfunc05(x, y int, m, n bool) int {
	return x + y
}

//可变长参数
func Myfunc06(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y) //类型是切片
}

func main() {
	fmt.Println(Myfunc01(10, 20))
	Myfunc02(15, 20)
	fmt.Println(Myfunc03())
	_, n := Myfunc04()
	fmt.Println(n)

	Myfunc06("长安忆", 1, 1, 1) //打印出来的是切片
}

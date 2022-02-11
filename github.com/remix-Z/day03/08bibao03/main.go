package main

import (
	"fmt"
	"strings"
)

//闭包简单例子1
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test01() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test.jpg  "))
	fmt.Println(txtFunc("test"))
}

//闭包简单例子2
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func test02() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11,9
	fmt.Println(f1(3), f2(4)) //12,8
	fmt.Println(f1(5), f2(6)) //13,7
}

func main() {
	fmt.Println("test01...")
	test01()

	fmt.Println("test02...")
	test02()

}

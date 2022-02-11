package main

import (
	"fmt"
)

//自定义类型
//type后面跟类型
type myInt int

//别名
type yourInt = int

func main() {
	var m myInt
	var y yourInt

	m = 100
	y = 100
	fmt.Printf("%T %T\n", m, y)

	var c rune //底层是type rune = int32
	c = '中'
	fmt.Printf("%T\n", c)
}

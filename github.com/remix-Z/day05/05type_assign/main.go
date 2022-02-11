package main

import (
	"fmt"
)

//类型断言
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错类型")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

//改良版本
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("字符串:", t)
	case int:
		fmt.Println("整形:", t)
	case bool:
		fmt.Println("波尔:", t)
	}

}

func main() {
	//assign(100)
	assign2(true)
	assign2("hello world")
}

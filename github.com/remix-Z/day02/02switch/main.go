package main

import "fmt"

func main() {
	var n = 5
	//switch n:=5;n{
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("none")
	}

	//fallthrough 和 goto 不推荐使用

}

package main

import "fmt"

const (
	red    int = 4
	yellow int = 2
	blue   int = 1
)

//二进制使用
//1 1 1
//红 100
//黄 010
//蓝 001
func color(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	color(red | yellow | blue)
}

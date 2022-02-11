package main

import (
	"bufio"
	"fmt"
	"os"
)

//输入时不能使用空格
func useScan() {
	var s string
	fmt.Println("请输入:")
	fmt.Scanln(&s)
	fmt.Println(s)
}

func useBufio() {
	var s string
	//bufio.NewReader()参数是接口类型 随便传
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}

func main() {
	useScan()
	//useBufio()
}

package main

import (
	"fmt"
	"strconv"
)

//go语言强类型

func main() {
	str := "10000"
	//i := int64(str)	不可以
	fmt.Println(str)

	//这样可以 强制类型转换
	i := int32(2316)
	ret1 := int64(i)
	//ret2 := string(i)	//不行
	ret2New := fmt.Sprintf("%d", i) //这样可以吧i转成字符串类型
	fmt.Println(ret1, ret2New)

	//把字符串转成int64类型
	ret3, err := strconv.ParseInt(str, 10, 64) //10:十进制 64:64位
	if err != nil {
		return
	}
	fmt.Printf("%T\n", ret3) //ret3 int64类型

	//转成int
	ret4, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	fmt.Printf("%T\n", ret4) //ret4 int类型

	a := 64
	ret5 := strconv.Itoa(a)
	fmt.Printf("%T\n", ret5)

	//从字符串解析出bool值
	boolStr := "false"
	ret6, err := strconv.ParseBool(boolStr)
	if err != nil {
		return
	}
	fmt.Printf("%T\n", ret6)
}

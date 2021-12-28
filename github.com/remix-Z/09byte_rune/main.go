package main

import (
	"fmt"
)

func main() {
	// 	a := 10

	// 	s2 := "白萝卜"
	// 	s3 := []rune(s2) //把字符串强制切片
	// 	s3[0] = '红'
	// 	fmt.Println(string(s3)) //把s3切片强制转换成字符串

	// 	c1 := "红"
	// 	c2 := '红'

	// 	fmt.Printf("c1:%T  c2:%T\n", c1, c2)

	// 	//类型转换
	// 	var f = float64(a)
	// 	//f = float64(a)
	// 	fmt.Println(f)
	// 	fmt.Printf("f:%T\n", f)
	//

	a := 10
	var b = float64(1.34)
	c := true
	d := "你好啊世界"

	//也可以用%v来输出变量的值
	fmt.Printf("a:[%T] [%d]\n", a, a)
	fmt.Printf("b:[%T] [%f]\n", b, b)
	fmt.Printf("c:[%T] [%t]\n", c, c) //%t输出bool类型的值
	fmt.Printf("d:[%T] [%s]\n", d, d)

	s1 := "hello沙河小王子"
	// s2 := "abcdefghijklmnopqrstuvwxyz"
	//这种写法打印出来的都是ASCII码 有问题
	// len := len(s2)
	// for i := 0; i < len; i++ {
	// 	fmt.Printf("%v\n", s2[i])
	// }

	//超纲部分 也不算 讨巧的写法
	//s2 := []rune(s1)
	sum := 0
	for _, c := range s1 {
		//也是根据ASCII来的，但实际汉字怎么排除
		if c > 122 {
			sum++
			//fmt.Printf("%c\n", c)
		}
	}
	fmt.Println(sum)

	fmt.Println(len(s1))
	fmt.Println(len([]rune(s1)))
	//计算差值 除以2 就是多出来汉字的个数，每个汉字占3字节
	n := 0
	n = len(s1) - len([]rune(s1))
	fmt.Println(n / 2)
}

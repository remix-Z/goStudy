package main

import "fmt"

//运算符
func main() {

	//+ - * / %
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //好像自动变成整型了
	fmt.Println(a % b)

	a++ //单独的语句，不能放在=号的右边赋值
	b--
	//i := a++ //nope
	fmt.Println(a)

	//关系运算符
	fmt.Println(a == b) //相同类型
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	//逻辑运算符
	age := 19
	if age > 18 && age < 60 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	if age < 18 || age > 60 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	//not取反
	isTrue := false
	fmt.Println(isTrue)
	fmt.Println(!isTrue)

	//位运算
	fmt.Println("**************")
	//按位与 	两位都为1则为1
	fmt.Println(5 & 2)
	//按位或	两位有一个为1则为1
	fmt.Println(5 | 2)
	//异或 ^ 	两位不一样则为1
	fmt.Println(5 ^ 2)
	//左移右移
	fmt.Println(5 >> 2) //101 -> 1
	fmt.Println(5 << 2) //101 -> 10100

	//var i int8 = 8
	//fmt.Println(i << 10)	会提示有问题

	//赋值运算符
	var x int
	x = 10
	x += 2
	x -= 2
	x /= 2
	x *= 2
	x %= 2
	x <<= 1
	x >>= 1
	x &= 1
	x |= 1
	x ^= 1
}

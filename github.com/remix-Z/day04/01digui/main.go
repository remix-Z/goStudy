package main

import "fmt"

//递归:自己调用自己
//但一定要有一个明确的退出条件
func main() {
	//阶乘
	ret := f1(5)
	fmt.Println(ret)

	//上台阶
	fmt.Println(f2(3))
}

//阶乘
func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

//斐波那契数列
//自己查
func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

//上台阶面试题
//n个台阶 一次可以走一步，也可以走两步，有多少种走法
func f2(n int) int {
	if 1 == n {
		return 1
	} else if 2 == n {
		return 2
	}

	return f2(n-1) + f2(n-2)
}

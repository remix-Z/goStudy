package main

import (
	"fmt"
)

//defer
func deferDemo() {
	fmt.Println("start")
	fmt.Println("hello world")
	defer fmt.Println("hey you") //defer把他后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("2222")
	defer fmt.Println("1111")
	fmt.Println("end")

	//defer用处
	//打开文件后需要关闭文件，为了防止忘记可以加一个defer语句
	//多用于函数结束前释放资源 文件句柄等等

}

//defer执行时机
//go语言中的return不是原子操作
//retrun x   ---> 先给返回值赋值=x,再执行返回
//return x 有defer ---> 先给返回值赋值=x,运行defer，再执行返回
//例
func example01() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//注释 先x=5，再x++ 再赋值
func example02() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// y=x=5,x++,y=5
func example03() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//内部的函数的x和example04函数的返回值不是一个x，所以还是5
func example04() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//指针传递
func example05() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
}

func main() {
	deferDemo()
	// 运行结果
	// start
	// hello world
	// end
	// 1111	多个defer 后进先出的规律
	// 2222
	// hey you

	fmt.Println(example01()) //5
	fmt.Println(example02()) //6
	fmt.Println(example03()) //5
	fmt.Println(example04()) //5
	fmt.Println(example05()) //6

}

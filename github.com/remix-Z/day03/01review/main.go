package main

import (
	"fmt"
)

//数组是值类型
func modify(a [3]int) {
	//go语言中函数传递的都是值 就相当于把一个文件从一个文件夹拿到另一个文件夹，更改的只是另一个文件夹的该文件，原文件夹中的该文件并没有更改
	a[1] = 100 //此处只是修改副本的值
}

func test01() {

	//运算符

	//数组
	//数组包含元素的类型和元素的个数。元素的个数(数组的长度)属于数据类型的一部分	跟C不一样
	// var ages1 [20]int 和 var ages2 [30]int
	//二者虽然数据类型一样，当长度不一样 所以不是一种数据类型

	//数组是值类型 ?怎么理解 跟c应该一样
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	modify(x)
	fmt.Println(x)

	//数组初始化方式 3种
	//1	跟2一样，建议用2
	var ages01 [30]int
	ages01 = [30]int{1, 2, 3, 4, 5}
	fmt.Println(ages01)
	//2
	var ages02 = [30]int{1, 2, 3}
	fmt.Println(ages02)
	//3
	var cities = [...]string{"上海", "沈阳"}
	fmt.Println(cities)
	//4
	var ages03 = [...]int{1: 2, 5: 10}
	fmt.Println(ages03)

	//二维数组
	//var a1 [3][2]int //外层3 内层2
	//多维数组只有最外层定义时可以使用...
	var a1 = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

}

func test02() {

	//切片 slice	引用类型(相当于一个框，取圈底层数组)
	//三要素 指针，长度，容量
	//定义 var s1 []int 声明的时候并未分配内存 需要先初始化再使用
	var s1 = []int{1, 2, 3}
	fmt.Println(s1)

	//make初始化
	var s2 = make([]bool, 2)
	fmt.Println(s2)
	var s3 = make([]bool, 0, 4) //初始化了分配内存了，只是长度为0，所以不是nil
	fmt.Println(s3)
	fmt.Println(s3 == nil) //false	虽然打印出来是空，但是是有内存空间的

	//切片不能比较 只能和nil

	//赋值拷贝 copy
	ss1 := []int{1, 2, 3}
	ss2 := ss1
	fmt.Println(ss2)
	ss2[1] = 100     //切片不存值，ss2和ss1一样都指向底层数组
	fmt.Println(ss1) //[1 100 3]
	fmt.Println(ss2) //[1 100 3]

	//切片扩容策略	***面试可能会问 建议了解
	//如果小于1024，那么直接扩两倍
	//如果大于1024，则1.25倍扩容
	//特殊情况见下函数
	test_kuorong()

}

//扩容策略
func test_kuorong() {
	a := [...]int{1, 2, 3, 4}
	s1 := a[:3]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("%p %p\n", s1, &a)
	s1 = append(s1, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000) //内存规格

	fmt.Println(s1, len(s1), cap(s1)) //容量是10
	fmt.Println(a)
	fmt.Printf("%p %p\n", s1, &a) //扩容后的s1指向了另一个底层数组 跟a已经没有关系了
	//那么问题又回来了 为什么是10
	//s1 容量是4 增加元素 4*2 < 3+6

	//int8
	array := []int8{} //空切片
	for i := 0; i < 16; i++ {
		array = append(array, 1, 2, 3, 4, 5) //增加2n-1个元素
		fmt.Print(cap(array), " ")
	}
	//8 16 16 32 32 32 64 64 64 64 64 64 128 128 128 128

	//int16
	fmt.Println()
	brray := []int16{}
	for i := 0; i < 16; i++ {
		brray = append(brray, 1, 2, 3, 4, 5)
		fmt.Print(cap(brray), " ")
	}
	//8 16 16 32 32 32 64 64 64 64 64 64 128 128 128 128

	//string
	fmt.Println()
	f := []string{}
	for i := 0; i < 16; i++ {
		f = append(f, "1.1", "2.2", "3.3", "4.4", "5.5")
		fmt.Print(cap(f), " ")
	}
	//5 10 20 20 40 40 40 40 80 80 80 80 80 80 80 80

}

//指针
func test03() {
	addr := "你好世界"
	addrP := &addr
	fmt.Println(addrP)        //内存地址
	fmt.Printf("%T\n", addrP) //*string类型

	fmt.Println(*addrP) //解引用
}

//map
func test04() {
	//var m1 map[string]int

	m1 := make(map[string]int, 10)
	m1["李明"] = 100
	m1["张三"] = 98
	m1["小华"] = 86
	m1["王柳"] = 59
	fmt.Println(m1)
	//如果key不存在，返回对应类型value的零值
	//fmt.Println(m1["小明"]) //0
	//为了防止上述情况，用下面的方法输出
	score, ok := m1["小明"]
	if !ok {
		fmt.Println("没有这个人")
	} else {
		fmt.Printf("分数是%d\n", score)
	}

	//注意点,即便delete了m1的全部key value，m1也不是nil，只有没有开辟空间才会为nil
	delete(m1, "小明")
	delete(m1, "张三") //如果delete的key是不存在的，就啥也不干不会报错
	fmt.Println(m1)

	//遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

}

//函数	可变参数
func test05(name string, x ...int) int {
	fmt.Printf("%T\n", x) //这里参数x是切片类型
	fmt.Println(x)
	return 1
}

//多个返回值
func f1(a int, b int) (sub int, sum int) {
	sum = a + b
	sub = a - b
	return //命名的返回值 return后可以省略
}

//在一个命名的函数内部不能声明另一个命名函数 但可以声明匿名的=。= 后面讲
func main() {

	//数组
	test01()
	//切片
	test02()

	test03()

	test04()

	test05("123", 1, 2, 3)
	//test05(" ", 1)
	fmt.Println(f1(2, 3))
}

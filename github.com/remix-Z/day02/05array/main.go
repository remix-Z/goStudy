package main

import "fmt"

//数组
//必须指定存放元素的类型和容量(长度，数组大小)
func main() {

	//可以改变成员，但是不能改变长度
	//数组的长度和类型都是数组的一部分，不同长度类型都是不能比较的
	var a1 [4]bool
	var a2 [5]int

	fmt.Printf("a1:%T,a2:%T\n", a1, a2)

	//初始化
	fmt.Println(a1, a2)
	a1 = [4]bool{true, true, true, true}
	//根据初始值自动推断数组的长度
	a100 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a100)
	a2 = [5]int{1, 2}
	fmt.Println(a2)
	a3 := [5]int{0: 1, 4: 2} //根据索引初始化
	fmt.Println(a3)

	//数组遍历
	cities := [5]string{"北京", "上海", "深圳", "广东", "成都"}
	//根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	//for range
	for i, v := range cities {
		fmt.Println(i, v)
	}

	//多维数组
	var a32 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	//fmt.Println(a32)
	//多维数组遍历
	for _, v1 := range a32 {
		//fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println("**********", b1, b2)

	//练习题
	//求数组{1,3,5,7,8}所有元素的和
	sum := 0
	array := [...]int{1, 3, 5, 7, 8}
	// for i := 0; i < len(array); i++ {
	// 	sum += array[i]
	// }
	//简化写法
	for _, v := range array {
		sum += v
	}
	fmt.Println(sum)

	//找出数组中和为指定值的两个元素的下标，比如从数组{1,3,5,7,8}中找到和为8的两个元素的下标分别为(0,3)和(1,2)
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

}

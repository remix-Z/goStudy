package main

import (
	"fmt"
	"strings"
	"unicode"
)

func test01() {

	//统计字符串中每个单词出现的次数，例如"how do you do"中how=1 do=2 you=1
	//s1 := "上海自来水来自海上"	只能判断英文 汉字就不行了
	s1 := "how do you do"
	s2 := strings.Split(string(s1), " ")
	sMap := make(map[string]int, len(s1))
	for _, r := range s2 {
		sMap[r]++
	}

	fmt.Println(sMap)
}

//判断字符串中汉字的个数
func test02() {
	//
	s1 := "vcghai自来水来自海上"
	var count1, count2 int
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count1++
		} else {
			count2++
		}
	}
	fmt.Printf("汉字的个数为%d\n", count1)
	fmt.Printf("非汉字的个数为%d\n", count2)
}

//回文判断
func test03() {
	//汉字的话就是占3字节，那么s1的len是9*3 = 27
	//s1 := "上海自来水来自海上"
	s1 := "上海自来水来自海上是"
	sp := []rune(s1)
	length := len(sp)

	for i := 0; i < length/2; i++ {
		if sp[i] != sp[len(sp)-1-i] {
			fmt.Println("不是回文字符串")
			return
		}
	}
	fmt.Println("是回文字符串")
}

func main() {
	test01()

	test02()

	test03()
}

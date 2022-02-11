package main

import (
	"fmt"
	"strings"
)

/*
	有50枚金币 需要分配给以下几人:Matthew Sarah Augustus Heidi Emilie Petter Giana Adriano Aaron Elizabeth
	分配规则如下:
	a. 名字中每包含1个'e'或'E'分1枚金币
	b. 名字中每包含1个'i'或'I'分2枚金币
	c. 名字中每包含1个'o'或'O'分3枚金币
	d. 名字中每包含1个'u'或'U'分4枚金币
	写一个程序 计算每个用户分到多少金币 以及最后剩余多少就hi比
	程序名 'dispatchCoin'
*/
// func dispatchCoin(m *map[string]int) {
// 	for k := range *m {
// 		// strings.Contains表示字符串是否包含
// 		if strings.Contains(k, "e") || strings.Contains(k, "E") {
// 			(*m)[k] += 1*strings.Count(k, "e") + 1*strings.Count(k, "E")
// 		}

// 		if strings.Contains(k, "i") || strings.Contains(k, "I") {
// 			(*m)[k] += 2*strings.Count(k, "i") + 2*strings.Count(k, "I")
// 		}

// 		if strings.Contains(k, "o") || strings.Contains(k, "O") {
// 			(*m)[k] += 3*strings.Count(k, "o") + 3*strings.Count(k, "O")
// 		}

// 		if strings.Contains(k, "u") || strings.Contains(k, "U") {
// 			(*m)[k] += 4*strings.Count(k, "u") + 4*strings.Count(k, "U")
// 		}
// 	}
// }

//分金币测试版
func dispatchCoinTest(m *map[string]int) {
	for k := range *m {
		// strings.Contains表示字符串是否包含
		(*m)[k] += 1*strings.Count(k, "e") + 1*strings.Count(k, "E")
		(*m)[k] += 2*strings.Count(k, "i") + 2*strings.Count(k, "I")
		(*m)[k] += 3*strings.Count(k, "o") + 3*strings.Count(k, "O")
		(*m)[k] += 4*strings.Count(k, "u") + 4*strings.Count(k, "U")
	}
}

func test01() {
	m1 := make(map[string]int, 10)

	m1["Matthew"] = 0
	m1["Sarah"] = 0
	m1["Augustus"] = 0
	m1["Heidi"] = 0
	m1["Emilie"] = 0
	m1["Petter"] = 0
	m1["Giana"] = 0
	m1["Adriano"] = 0
	m1["Aaron"] = 0
	m1["Elizabeth"] = 0

	dispatchCoinTest(&m1)
	//fmt.Println(m1)

	sum := 50
	for k, v := range m1 {
		fmt.Println(k, v)
		sum -= v
	}
	fmt.Printf("剩余金币数为%d\n", sum)

}

//封装 方便日后修改
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Petter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	//fmt.Println(len(users))
	//遍历用户
	for i := range users {
		distribution[users[i]] += 1*strings.Count(users[i], "e") + 1*strings.Count(users[i], "E")
		distribution[users[i]] += 2*strings.Count(users[i], "i") + 2*strings.Count(users[i], "I")
		distribution[users[i]] += 3*strings.Count(users[i], "o") + 3*strings.Count(users[i], "O")
		distribution[users[i]] += 4*strings.Count(users[i], "u") + 4*strings.Count(users[i], "U")
	}

	for k, v := range distribution {
		fmt.Println(k, v)
		coins -= v
	}

	return coins

}

func main() {
	left := dispatchCoin()
	fmt.Printf("剩余金币数:%d\n", left)
}

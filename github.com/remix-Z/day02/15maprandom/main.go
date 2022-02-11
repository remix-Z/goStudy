package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano()) //UnixNano 纳秒

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头字符串
		value := rand.Intn(100)          //随机0-99整数
		scoreMap[key] = value
	}

	fmt.Println(scoreMap)
	fmt.Printf("*******************\n")
	//切片
	keys := make([]string, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

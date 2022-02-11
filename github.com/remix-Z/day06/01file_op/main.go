package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//坑:os.O_APPEND只能在文件末尾追加
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileObj.Close()

	//func (f *File) Seek(offset int64, whence int) (ret int64, err error)
	fileObj.Seek(2, 0) //光标移动	//a/r/n b/r/n 移动三个位置才是换行 b

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(ret[:n]))

	//这种写法是按位替换后面的字符
	c := []byte{'C'}
	fileObj.Write(c)
	
	//bufio
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("你好")
	writer.Flush()

	fileObj.WriteString("中国")
}

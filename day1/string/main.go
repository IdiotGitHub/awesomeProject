package main

import (
	"fmt"
	"strconv"
)

/*
字符串常用的系统函数
	1.统计字符串的长度（注意返回的是字节长度，而不是实际字符串的字符个数，但是如果传入的是一个数组或者切片的话返回的就是个数。ASCII占一个字节，汉字占3个字节），按字节len(str)-->它是一个内建函数直接使用即可
	2.将字符串转为切片[]rune(str)
	3.字符串转整数num, err := strconv.Atoi(str)
	4.整数转字符串str := strconv.Itoa(num)
	5.字符串转byte切片 []byte  var bytes = []byte(str)
	6.[]byte转字符串 var str = string([]byte{97, 98, 99})
*/
func main() {
	lenDemo("hello北京")
}

//len()demo
func lenDemo(str string) {
	fmt.Println("str's len =", len(str))
	//这种方式是不可取的，因为len()取的是字节数
	for i := 0; i < len(str); i++ {
		fmt.Printf("字符%d=%c\n", i, str[i])
	}
	//如果想遍历含有中文的字符串时，可以先将字符串转为切片再使用len()进行遍历
}
func runeDemo(str string) {
	strs := []rune(str)
	for i := 0; i < len(strs); i++ {
		fmt.Printf("字符%d=%c\n", i, strs[i])
	}
}

func AtoiDemo(str string) {
	fmt.Println(strconv.Atoi(str))
}

func ItoaDemo(num int) {
	fmt.Println(strconv.Itoa(num))
}

func byteSliceDemo(str string) {
	var bytes = []byte(str)
	fmt.Printf("bytes=%v\n", bytes)
}

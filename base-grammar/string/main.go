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
	7.10进制数字转2、8、16进制数使用strconv.FormatInt(num, 2、8、16)
	8.查找子串是否存在指定的字符串：strings.Contains(str1, substr)返回bool值
	9.统计字符串中有几个指定字串，strings.Count("abcddd", "d")返回3
	10.不区分大小写的字符串比较（==是区分字母大小写的）：strings.EqualFold(str1, str2)
	11.返回子串在字符串中第一次出现的索引位置（从0开始看）如果子串不在字符串中返回-1
	12.返回子串在字符串中最后一次次出现的索引位置（从0开始看）如果子串不在字符串中返回-1
	13.将指定的子串替换成另一个子串：strings.Replace("go go go hello", "go", "golang", n),n可以指定你希望替换几个，-1表示全部替换，它是生成一个新的字符串，源字符串没有发生变化
	14.按照指定的某个字符为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok", ",")
	15.将字符串进行大小写转换strings.ToLower(str),strings.ToUpper(str)
	16.将字符串左右两侧的空格去掉， string.TrimSpace(str)
	17.将字符串左右两边指定的字符去掉，strings.Trim("! hello! ", " !")将左右两边都空格和！去掉
	18.将字符串左边指定的字符去掉，strings.TrimLeft("! hello! ", " !")将左边都空格和！去掉
	19.将字符串右边指定的字符去掉，strings.TrimRight("! hello! ", " !")将右边都空格和！去掉
	20.判断字符串是以指定的字符串开头：strings.HasPrefix(str1,str2)返回bool
	21.判断字符串是以指定的字符串结尾：strings.HasSuffix(str1,str2)返回bool


*/
func main() {
	//lenDemo("hello北京")
	AtoiDemo("hello")
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
	num, err := strconv.Atoi(str)
	//当转换出错时，num不会被赋予值，而是被初始化默认值
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}
}

func ItoaDemo(num int) {
	fmt.Println(strconv.Itoa(num))
}

func byteSliceDemo(str string) {
	var bytes = []byte(str)
	fmt.Printf("bytes=%v\n", bytes)
}

//[]byte 转string
func byteSliceToString(bytes []byte) {
	fmt.Println(string(bytes))
}

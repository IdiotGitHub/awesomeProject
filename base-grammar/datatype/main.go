package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

//go中没有double类型，使用float64来代替
//go语言中同样没有char类型，使用byte代替

/**
数据类型-
		-->基本数据类型
			-->1.数值型
				-->1.整数类型
					-->int也可以存字符！（由于byte类型范围是0~255，所以只能存储ASCII码的字符，如果超出了ASCII就可以使用int或uint来存储）
					-->int，默认是有符号的，根据系统来判单其长度，如果系统是32位的那么int表示int32，如果是64位的那么就表示int64
					-->无符号整数(uint8=0~255,uint16=0~2e*16-1,uint32=0~2e*32-1,uint64=0~2e*64-1)
				-->2.浮点类型,都是有符号的
			-->2.字符型 byte其实就相当于uint8,
				-->如果直接输出byte类型变量，那么将会输出数字，相当于输出uint8，想要输出字符那么需要使用fmt.Printf("%c",var)来输出
				-->字符类型是可以进行运算的，相当于一个整数
			-->3.布尔型bool,占一个字节，只能是true或者false
			-->4.字符串
			-->5基本数据类型的默认值：
				-->1.整型：0
				-->2.浮点型：0
				-->3.字符串：""
				-->4.布尔值：false
		-->派生/复杂数据类型
			-->1.指针
				-->
			-->2.数组
			-->3.结构体
			-->4.管道
			-->5.函数
			-->6.切片
			-->7.接口
			-->8.map

	-->基本数据类型的相互转换
		-->go语言中的基本数据类型之间的转换需要显示转换，不能够自动转换
		--> T(v):int(0.2)等等
*/
func main() {
	//stringFunc()
	//defaultValue()
	//convertStringToBasicType()
	pointer()
}

//整型
func intFunc() {
	var i = 65535
	//使用fmt.Printf()进行格式化输出，并使用%T来输出类型
	fmt.Printf("i 的数据类型是%T\n", i)
	//输出某个变量所占用的字节大小
	fmt.Printf("i的数据类型是：%T,i所占用的字节大小为：%d\n", i, unsafe.Sizeof(i))
	var num float64 = 1
	fmt.Printf("num=%f\n", num)
	var char int = '你'
	fmt.Printf("char = %c\n", char)
	fmt.Printf("char = %d\n", char+i)
}

//布尔类型
func boolFunc() {
	var b = true
	fmt.Printf("i的数据类型是：%T,i所占用的字节大小为：%d\n", b, unsafe.Sizeof(b))
}

//字符串
func stringFunc() {
	//string的基本使用
	//双引号可以识别转义字符，反单引号可以原样输出字符串
	var address string = "北京长城"
	fmt.Println(address)
	str1 := ``
	str1 += `hello\n world`
	fmt.Println(str1)
	//如果go语言中的字符串特别长，可以分行表示，但是在使用加号操作符的时候需要将加号放在末尾
	str2 := "hello" + "world" +
		"hello" + "world"
	fmt.Println(str2)
}

//基本数据类型的默认值
func defaultValue() {
	var num int
	var flo float32
	var boo bool
	var str string
	fmt.Printf("num=%v,flo = %v,boo = %v, str = %v", num, flo, boo, str)
}

//数据类型转换
func convertType() {
	var num int = 100
	fmt.Println(float32(num))
	var flo float32 = float32(num)
	fmt.Println(flo)
}

//基本数据类型和string类型之间的转换
func convertBasicTypeToString() {
	//1.使用fmt包中的Sprintf("%参数",表达式)
	var num1 int = 10
	var num2 float32 = 2.34
	var boo bool = true
	fmt.Println(fmt.Sprintf("%d", num1))
	fmt.Println(fmt.Sprintf("%f", num2))
	fmt.Println(fmt.Sprintf("%t", boo))
	//2.使用strconv包中的几个函数，比较麻烦，尽量使用fmt.Sprintf()。
	//-->strconv.FormatInt(num1,10)
	//-->strconv.FormatFloat(num2,'f',10,64)
	//-->strconv.FormatBool(boo,10)
	//-->strconv.Itoa()

}
func ConvertStringToBasicType() {
	//使用strconv包中的几个函数
	var str1 string = "hello"
	var boo bool
	//这里肯定会转换失败，那么即使boo原来有值也会被赋值为默认值
	boo, _ = strconv.ParseBool(str1)
	fmt.Println(boo)
	var str2 = "123"
	//这个地方不太好，在进行类型转换的时候如果不是要int64的话，还需要一个中间值来转，有点尬
	var num1 int64
	num1, _ = strconv.ParseInt(str2, 10, 0)
	fmt.Println(num1)
	//float也一样
}

func pointer() {
	//基本数据类型在内存中的分配
	var num1 int = 1
	fmt.Println("num1的地址为：", &num1)
	//指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	//指针类型跟基本数据类型的区别就是，指针类型存的是地址，它本身也是内存中的一块，也有自己的内存地址
	//它所能接受的只能是地址
	mid := int64(num1)
	var ptr *int64 = &mid
	fmt.Println(&ptr)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	//注意，值类型包括：基本数据类型int系列，float系列，bool、string、数组和结构体
	//-->值类型一般在栈内分配内存
	//引用类型：指针、slice切片、map、管道chan、interface等
	//-->引用类型一般在堆内分配内存

}

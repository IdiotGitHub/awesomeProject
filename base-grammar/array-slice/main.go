package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
数组与切片
	数组可以存放多个同一类型数据，数组也是一种数据类型，在go中，数组是值类型。
	数组的首地址就是数组第一个元素的地址；可以通过&数组名来获取数组地址，第二个元素的地址就是首地址加该元素的大小，后面的元素地址同理
	1.四种初始化数组方式：
		1.声明时直接初始化
			var arr1 [3]int = [3]int{1, 2, 3}
		2.声明时直接初始化（使用类型推导的方式）
			var arr1 = [3]int{1, 2, 3}
		3.数组长度可以使用系统判断的方式
			var arr1 = [...]int{1, 2, 3}
		4.在赋值的时候指定下标
			var arr1 [3]int = [3]int{0:1, 2:2}
	2.数组的遍历
		1.普通的for循环即可
			for i:=0;i < len(arr1); i++ {}
		2.使用for--range结构进行遍历
			for index, value := range array {}
			第一个返回值index是数组下标
			第二个value是在该下标位置的值
			他们都是仅在for循环内部可见的局部变量
			遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线
			index和value的名称不是固定的 ，程序员可以自行指定，一般命名为index，value
	3.数组使用的注意事项和细节
		1.数组是多个相同类型的数据组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
		2.var arr []int 这时arr就是一个slice切片
		3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
		4.数组创建之后不进行初始化赋值，系统会自动分配默认值（具体类型的默认值同变量的默认值）
		5.使用数组的步骤：1.声明数组并开辟空间，2.给数组各个元素赋值，3.使用数组
		6.数组的下标是从0开始的
		7.数组下标必须在指定范围内使用，否则报panic，数组越界
		8.go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
		9.如果想在其他函数中去修改原来的数组，可以使用引用传递（指针方式）
			注意：使用引用传递的时候参数格式为：arr *[2]int 而不是arr [2]*int。前者为引用传递，后者为值传递的int指针类型的参数。在函数内部进行访问时可以直接使用arr[0]的方式进行访问，也可以使用这个方式(*arr)[0]进行访问
		10.长度是数组类型的一部分，在传递函数参数时需要考虑数组的长度。
	**Go语言中的算术运算有点操蛋了，在进行除法运算的时候一定要注意int/int，在进行运算的时候一定要手动强转为float64
切片：
	1.切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
	2.切片的使用和数组类似，遍历、访问元素和求切片长度都和数组一样。
	3.切片的长度是可以变化的，因此切片是一个可以动态变化的数组。
	切片本质上其实是一个结构体type slice struct{ptr *[]type,len int, cap int}
	切片的使用
		1.定义一个切片，然后让切片去引用一个已经创建好的数组
		2.通过make来创建切片。var slice1 []int = make([]int, 3, 10)
						   var name []type = make([]type, len, cap);cap >= len
			通过make方式创建的切片指向的数组由make底层维护，对外观不可见，即只能通过slice去访问各个元素
		3.定义一个切片，直接就指定具体数组，使用原理类似make的方式。var slice []string = []string{"a", "b", "c"}
切片注意事项和细节说明
	1.切片初始化时var slice = arr[startIndex:endIndex]从arr数组下标为startIndex,取到endIndex的元素（前闭后开）
	2.切片在初始化的时候仍然不能越界，但是可以动态增长
		1.var slice = arr[0:end]可以简写为var slice = arr[:end]
		2.var slice = arr[start:len(arr)]可以简写var slice = arr[start:]
		3.var slice = arr[0:len(arr)]可以简写var slice = arr[:]
	3.cap是一个内置函数，用于同即切片的容量，即最大可以存放多少元素
	4.切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间共切片使用
	5.切片可以继续切片
	6.用append内置函数动态扩充切片
		var slice1 []int = []int{1, 2, 3, 4}
		slice1 := append(slice1, 5, 6, 7)
		slice2 := append(slice1, 5, 6, 7)
	也可以使用append(slice,slice...)参数2只能时切片不能是数组，也不能遗漏...
	切片append操作的底层原理：
		1.切片append操作的本质就是对数组扩容
		2.go底层会创建一个新的数组newArr
		3.将slice原来包含的元素拷贝到新的数组newArr
		4.slice重新引用到newArr
		5.注意newArr是在底层来维护的，程序员不可见
	7.切片的拷贝操作
		切片使用copy内置函数完成拷贝（要求两个参数都是切片）
slice和string
	1.string底层是一个byte数组，因此可以使用切片来截取字符串（其实跟其他语言相同）
	2.string是不可变的，因此不可以使用str[0] = "a"来改变字符串
	3.如果需要修改字符串，可以先将字符串转成一个切片，修改完成之后再转为string
		但是这种方式只能处理英文和数字，不能处理中文字符（中文字符太大了）
		将切片转成[]rune即可，它是按字符处理，兼容中文字符

*/

func main() {
	forRangeDemo()
	arr := [2]int{1, 2}
	arrayDemo(&arr)
	fmt.Println(arr)
	arrayExcise()
	arrayMax()
	reverseArray()
	sliceDemo2()
	sliceAndString()
}

//for-range
func forRangeDemo() {
	var arr1 [3]string
	for index, _ := range arr1 {
		arr1[index] = "Jim" + strconv.Itoa(index)
	}
	for _, value := range arr1 {
		fmt.Println(value)
	}
}

//数组引用传递
func arrayDemo(arr *[2]int) {
	arr[1] = 5
	(*arr)[0] = 3
}

//数组练习
func arrayExcise() {
	var arr [26]byte
	for i := 0; i < 26; i++ {
		arr[i] = byte(int('A') + i)
	}
	for _, value := range arr {
		fmt.Println(string(value))
	}
}

//求数组最大值的下标
func arrayMax() {
	var arr = [5]int{1, 3, 5, 12, 41}
	max := arr[0]
	index := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			max = arr[i+1]
			index = i + 1
		} else {
			max = arr[i]
			index = i
		}
	}
	fmt.Println("max=", max, "下标为", index)
}

//随机生成数组元素并 反转打印数组
func reverseArray() {
	var arr [5]int
	//注意在使用随机数生成时，要指定一个种子
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) //随机生成[0,100)的随机整数
	}
	fmt.Println("随机生成的数组为：", arr)
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])

	}
}
func sliceDemo() {
	var slice1 []float64 = make([]float64, 3, 10)
	fmt.Println(slice1)
}
func sliceDemo2() {
	var slice []int = []int{1, 2, 3, 4}
	fmt.Printf("%v\n", &slice[0])
	slice = append(slice, slice...)
	fmt.Printf("%v\n", &slice[0])
}
func sliceAndString() {
	str := "helloworld"
	slice := str[strings.Index(str, "o"):]
	fmt.Printf("slice=%v\n", slice)
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
*/

func main() {
	forRangeDemo()
	arr := [2]int{1, 2}
	arrayDemo(&arr)
	fmt.Println(arr)
	arrayExcise()
	arrayMax()
	reverseArray()
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

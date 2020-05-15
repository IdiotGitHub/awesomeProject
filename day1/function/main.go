package main

import (
	"awesomeProject/day1/function/utils"
	"fmt"
	"strconv"
	"strings"
)

/*
golang中的函数
	突然发现在使用float类型进行计算的时候会发生精度上的问题，造成结果错误。
	那么在使用的时候需要先将计算的结果进行转换。
	切记在使用float的时候会有精度问题，在使用float进行计算的时候结果必为错误的！！
	在进行浮点数类型比较的时候也最好是转成字符串进行比较
	先将float类型转为字符串，要保留多少位要根据实际情况进行判断。再将字符串转为float类型。真尼玛恶心
函数的调用机制
	1.基本数据类型进行函数参数传递的时候是值传递
	2.函数调用的时候会新开辟一个空间，用来存变量，编译器会通过自身的处理让这个新的空间和其他的栈空间区分开来
	3.在每个函数对应的栈中，数据空间是独立的，不会混肴
	4.当一个函数调用完毕后，程序会销毁这个函数对应的栈空间
return语句
	1.Go函数支持返回多个值
	2.如果返回多个值时，在接收时，希望忽略某个返回值，则使用_符号表示占位忽略
	3.如果返回值只有一个，返回值类型列表可以不写括号
函数的递归调用
	1.执行一个函数时，就创建一个新的受保护的独立空间
	2.函数的局部变量是独立的，不会相互影响
	3.递归必须向推出条件逼近，否则就会无限递归，死循环
	4.当一个函数执行完毕，或者遇到return，就会返回，谁调用，就将结果返回给谁；同时当函数执行完毕或者返回时，改函数本身也会被系统销毁
函数注意事项和细节
	1.函数的形参列表可以是多个，返回值列表也可以是多个。
	2.形参列表和返回值类型的数据类型可以是值类型和引用类型
	3.函数的命名遵守标识符命名规范，首字母不能是数字，首字母大写改函数可以被本包文件和其他包文件使用，类似Java中的public；首字母小写，只能被本包文件使用，类似Java中的private。

	4.函数中的变量是局部的，函数外不生效
	5.基本数据类型和数组默认都是值传递的，即进行值拷贝，在函数内修改，不会影响到原来的值。
	6.如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。类似引用变量传递
	7.Go函数不支持重载
	8.在Go中，函数也是一种数据类型，可以赋值给一个变量，该变量就是一个函数类型的变量。通过该变量可以对函数调用。(在进行赋值操作的时候只需要将函数名赋值给变量即可)
	9.函数既然是一种数据类型，因此在Go中，函数是可以作为形参的。其格式为
		func myFunc(funVar func(int, int) int, num1 int, num2 int) int {
			//函数体
			return funVar(num1, num2)
		}
	10.为了简化数据类型定义，Go语言支持自定义数据类型
		基本语法：type 自定义数据类型名 数据类型 //相当于起别名
		eg：type myFunc func(int, int) int //这时 myFunc就是 func(int, int) int类型
	11.Go支持对函数返回值命名，在命名之后可以对返回值进行修改，在函数结束时直接return即可
	12.使用_标识符忽略返回值。
	13.Go函数支持可变参数,而且必须将可变参数放在形参列表的最后一项
		func myFunc(args... int) int {} //args就是一个切片
Go中的init函数
	1.每一个源文件都可以包含一个init函数，改函数会在main函数执行前被条用
	2.如果一个文件同时包含变量定义，init函数和main函数，则执行的流程是全局变量定义->init函数->main函数
	3.如果引入的文件中包含init函数那么先执行被引入文件中的init函数
Go中匿名函数
	Go支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。
	匿名函数调用方式，1.在定义匿名函数时就直接调用，这种方式匿名函数只被执行一次 func (num1 int, num2 int) int {} (num1, num2)
					2.可以将匿名函数赋值给一个变量，那么该变量就是函数类型，可以被多次调用
					3.全局匿名函数，如果将匿名函数赋值给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，可以在程序中有效。
闭包
	闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）
	eg:
	func addUpper() func(int) int {
	var num int = 10
	return func(i int) int {
		num += i
		return num
		}
	}
	返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包。
	***闭包可以理解成一个类，将它保存在某一个变量中就相当于初始化，后面每次调用就不需要再进行初始化
defer
	在函数中，程序员经常需要创建资源（比如数据库连接、文件句柄、锁等），为了在函数执行完毕之后，及时的释放资源，Go的设计者提供了defer（延时机制）
	1.当执行到defer时会将defer后的语句压入一个独立的栈中
	2.当函数执行完毕后，再从defer栈中执行（后入先出顺序）
	3.在defer入栈的时候会将相关的值变量进行值拷贝一起入栈
	一般是创建资源后，可以立即defer将其关闭。然后后面可以继续使用该资源，不用考虑关闭资源的时机（或者忘记关闭资源）。
函数参数的传递方式
	两种传递方式
		1.值传递（值类型：基本数据类型int系列，float系列，bool，string，数组和结构体）
		2.引用传递（指针，slice切片，map，管道chan，interface）
		其实，不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是引用的拷贝。一般来说，地址拷贝效率高，因为数据量小，而值拷贝取决于拷贝的数据大小（比如结构体也是值拷贝，当结构体数据打的时候效率就会低）

*/
var age = testInit()

func testInit() int {
	fmt.Println("testInit()")
	return utils.Age
}
func init() {
	fmt.Println("init()")
}
func main() {
	/*	fmt.Println(2.2 == Float64(3.3-1.1))
		fmt.Println(getSumAndSub(11, 21))
		f := getSum
		fmt.Printf("f的类型为%T, getSum的类型为%T\n", f, getSum)
		fmt.Println(getSum2(1, 2, 3, 4, 5, 6))
		num1 := 2
		num2 := 3
		exercise1(&num1, &num2)
		fmt.Println("交换2和3的值", num1, num2)
		unknownFuncDemo2()*/
	f := biBao(".jpg")
	fmt.Println(f("world.jpg"))
	fmt.Println(f("world"))
	//时间和日期
	utils.Demo1()
	utils.Demo2()
	utils.Demo3()
}
func Float64(num float64) float64 {
	float64Num, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", num), 64)
	return float64Num
}
func getSumAndSub(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
func getSum(num1 int, num2 int) int {
	return num2 + num1
}
func getSum2(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

//交换两数的值
func exercise1(num1 *int, num2 *int) {
	*num1 = *num1 + *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2
}

//匿名函数调用方式1
func unknownFuncDemo1() {
	result := func(num1, num2 int) int {
		return num2 + num1
	}(10, 2)
	fmt.Println(result)
}

//匿名函数调用方式2
//可以将匿名函数赋值给一个变量，那么该变量就是函数类型，可以被多次调用
func unknownFuncDemo2() {
	result := func(num1, num2 int) int {
		return num2 + num1
	}
	fmt.Println(result(10, 20))
}

//闭包
//编写一个函数makeSuffix(suffix string) 可以接受一个文件后缀名（比如.jpg)。并返回一个闭包
//调用闭包，可以传入一个文件名，如果改文件名没有指定的后缀，则返回文件名.jpg。如果有后缀名返回文件名
//如果不使用闭包的话需要每次都传入这个后缀名

func biBao(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return s
		}
		return s + suffix
	}
}

//累加器
func addUpper(num int) func(int) int {
	return func(i int) int {
		num += i
		return num
	}
}

//defer
func deferDemo1() int {
	//当执行到defer时会将defer后的语句压入一个独立的栈中
	//当函数执行完毕后，再从defer栈中执行（后入先出顺序）
	defer fmt.Println("deferDemo->defer1")
	defer fmt.Println("deferDemo->defer2")
	res := 30
	fmt.Println(res)
	return res
}

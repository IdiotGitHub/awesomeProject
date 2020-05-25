package main

import "fmt"

func main() {
	//如果除号两边都是整数那么结果就一定是整数类型，想要浮点类型的结果的话就需要有浮点数参与运算。
	//一定要注意
	var num1 float64
	num1 = 10 / 4
	fmt.Printf("num1 = %v, type is %T\n", num1, num1)
	//取模的公式 a % b = a - a / b * b
	fmt.Println("10 % 3 = ", 10%3)     //1
	fmt.Println("-10 % 3 = ", -10%3)   //-1
	fmt.Println("10 % -3 = ", 10%-3)   //1
	fmt.Println("-10 % -3 = ", -10%-3) //-1
	//自加自减只能当作独立语句使用，不能用于赋值语句中，也不能用于条件判断中。
	//在go语言中只存在后加加 ，不存在前加加；减号同理
	//a := i++是错误的
	var i int64
	i++
	fmt.Println(i)
	//关系运算符，结果都是bool类型
	//关系运算符组成的表达式，称为关系表达式
	//逻辑运算符，将多个关系表达式连接起来，结果也为bool类型

	//在go语言中，如果要交换两个变量的值需要使用中间变量来过度。
	//不适用中间变量来交换两个变量的值
	var a = 10
	var b = 20
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}

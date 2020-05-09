package main

import "fmt"

/**
变量的作用域
	1.局部变量，函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部。
	2.全局变量，函数外部声明/定义的变量叫全局变量，作用域在整个包内都有效，如果其首字母为大写，则作用域在整个程序有效（使用该变量的话需要用包名加点使用）
		声明定义全局变量的时候不可以使用类型推导的方式，
	3.如果变量是在一个代码块中，比如if/for中，那么这个变量只在这个代码块中有效。
 */

func main() {
	var i int

	i = 12

	fmt.Println(i)

	//一次声明多个变量
	//1.
	var num1, name, num2 = 100, "tom", 23
	fmt.Print("num1 = ", num1, "num2 = ", num2, "name = ", name)
	//2.
	var (
		n1 int
		n2 string
		//小数的默认值也是0
		n3 float32
	)
	n1 = 1
	n2 = "hello"

	fmt.Print(n1, n2, n3)

}

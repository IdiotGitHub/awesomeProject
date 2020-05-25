package main

import (
	"encoding/json"
	"fmt"
)

/*
方法
	Golang中的方法是作用在指定的数据类型上的，因此自定义类型，都可以有方法，而不仅仅是struct
	方法只能通过与其绑定的对象调用，不可以直接调用，也不能被其它类型对象调用
方法的调用和传参机制
	方法的调用和传参机制与函数基本一致，不一样的地方是方法调用时，会将调用方法的结构体对象，当作实参传递给方法
方法的声明（定义）
	func (receiver type) methodName(参数列表) (返回值列表) {方法体 return 返回值}
	1.参数列表，表示方法输入
	2.receiver type表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
	3.receiver type，type可以是结构体也可以是其他的自定义类型
	4.receiver就是type类型的一个变量
	5.返回值列表，表示返回的值，可以是多个
	6.方法主体，表示为了实现某一功能的代码块
	7.return语句不是必须的
方法注意事项和使用细节
	1.结构体类型是值类型，遵循值类型的传递机制，是值拷贝传递方式
	2.如果程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
	3.Golang中的方法作用在指定的数据类型上。因此自定义类型，都可以由方法，而不仅仅是struct，比如int，float32都可以由方法
	4.方法的访问范围控制的规则，和函数一样，方法名字首字母小写，只能在本包内访问，方法首字母大写，可以在本包和其他包访问。
	5.如果一个类型实现了String()这个方法，那么fmt.Println()默认会调用这个变量的String()进行输出
		**注意，这里必须传结构体对象的地址，而且String()方法的方式必须是func (receiver type) String() string { method Body;return str}
方法和函数的区别
	1.调用方式不一样
		函数的调用方式，直接函数名调用
		方法的调用方式，需要结构体对象加点的方式进行调用
	2.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
		对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，返货来同样也可以（Go语言进行了优化处理）（就是当方法接收一个值类型的结构体变量的话，无论是值类型的结构体变量调用还是指针类型的结构体变量调用都不会改变结构体变量的值）
		如果是指针类型的接收者，在调用时无论是值类型的结构体变量还是指针类型的结构体变量进行调用都会改变结构体变量
		*********能不能改变结构体变量取决于方法的接收者是值类型还是指针类型
*/
type A struct {
	Num int
}

//要注意，结构体类型在传参的时候是值传递。想要改变结构体内部的值，需要使用地址传递
//其实在结构体结构体复杂后，使用的更多的是地址传递，效率更高
func (a A) modifyNum(num int) {
	a.Num = num
	fmt.Println(a.Num)
}
func (a *A) String() string {
	str := fmt.Sprintf("a-->%v", a.Num)
	return str

}
func (a *A) modifyNumAddress(num int) {
	(*a).Num = num
}
func main() {
	a := A{
		Num: 12,
	}
	//这里必须传入结构体对象的地址
	fmt.Println(&a)
	/*	a := A{
			Num: 12,
		}
		a.modifyNum(23)
		fmt.Println(a.Num)
		//这样也是可以的，Go语言底层会自动处理，两种使用方式是等价的
		a.modifyNumAddress(33)
		fmt.Println(a.Num)
		//标准的使用方式
		(&a).modifyNumAddress(38)
		fmt.Println(a.Num)*/
	circle := Circle{
		radius: 10.0,
	}
	fmt.Println(circle.area())
	m := MethodUtils{}
	m.print()
	m.printParams(3, 4)
	fmt.Println(m.area(5, 3))
	m.oddEven(3)
	//student test
	student := Student{
		"小许",
		1,
		23,
		99,
	}
	fmt.Println(student.say())
	//单行写最后可以不加逗号，多行写最后一个属性后面就必须加逗号
	//student2 := Student{"小许",1, 23, 99}
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14
}

//excise1
type MethodUtils struct {
}

func (methodUtils *MethodUtils) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (methodUtils *MethodUtils) printParams(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (methodUtils *MethodUtils) area(m, n int) int {
	return m * n
}

//判断奇偶
func (methodUtils *MethodUtils) oddEven(num int) {
	fmt.Println(num/2 == 0)
}
func (methodUtils *MethodUtils) calculator(m, n float64, operator string) {
	switch operator {
	case "+":
		fmt.Println("两个数的和为：", fmt.Sprintf("%.2f", m+n))
	case "-":
		fmt.Println("两个数的差为：", fmt.Sprintf("%.2f", m-n))
	case "*":
		fmt.Println("两个数的积为：", fmt.Sprintf("%.2f", m*n))
	case "/":
		fmt.Println("两个数的商为：", fmt.Sprintf("%.2f", m/n))
	default:
		fmt.Println("输入了错误的运算符")
	}
}

type Student struct {
	Name   string
	Gender int
	Age    int
	Score  float64
}

func (s *Student) say() string {
	//fmt.Println(*s)
	//这个地方无论是用s、&s、*s结果都是一样的，Golang进行了优化
	info, _ := json.Marshal(*s)
	return string(info)
	//也可以使用fmt.Sprinf()生成字符串
}

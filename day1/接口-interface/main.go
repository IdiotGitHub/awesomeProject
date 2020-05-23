package main

import "fmt"

/*
接口
	interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量，到某个自定义类型（struct）需要使用的时候，再根据具体情况把这些方法实现
	type interfaceName interface{
		method1(params type) returnList
		method2(params type) returnList
		...
	}
	接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想。
	Golang中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口，
接口注意事项和使用细节
	1.接口本身不能创建实例，但是他可以指向实现了该接口的自定义类型的变量
	2.接口中的所有方法都没有方法体
	3.Golang中，一个自定义类型需要将某个接口的所有方法都是实现，才是该自定义类型实现了这个接口
	4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口变量
	5.只要是自定义数据类型就可以实现接口，不仅仅是结构体类型
	6.一个自定义类型可以实现多个接口
	7.Golang接口中不能有任何字段
	8.一个接口可以继承多个接口，但是不允许继承有相同方法的接口
	9.interface类型默认是一个指针类型，如果没有对interface初始化就是用，那么会输出nil
	10.空接口interface{}，没有任何方法，所有所有类型都实现了空接口，因此可以把任何一个变量赋给一个空接口
*/
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//实现接口的时候用的是指针类型就传地址过去，用的是值类型就直接传过去就行
	computer.Working(&phone)
	computer.Working(&camera)

	//接口可以指向实现了该接口的自定义类型变量
	a := A{
		1,
	}
	//这个地方同样也是，在a实现Interface的时候使用的是指针类型就需要把a的地址传递过去
	var in Interface = &a
	in.Test()
	//当D接口继承B和C的时候，只要自定义类型A实现了D，那么同时也实现了B和C
	var b B = &a
	var c C = &a
	var d D = &a
	b.test1()
	c.test2()
	d.test1()
}

type Usb interface {
	Start()
	Stop()
}
type Computer struct {
}

//通过参数是接口进行传参，可以实现多态，这时候就必须传入一个实现了该接口的自定义变量（struct）
func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("Phone is starting...")
}
func (p *Phone) Stop() {
	fmt.Println("Phone is stopped...")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("Camera is starting...")
}
func (c *Camera) Stop() {
	fmt.Println("Camera is stopped...")
}

//接口不可以直接创建实例，但是可以指向实现了该接口的自定义类型变量
type Interface interface {
	Test()
}

//自定义类型A
type A struct {
	Id int
}

//自定义类型A实现接口Interface
func (a *A) Test() {
	fmt.Println((*a).Id)
}

//一个接口可以实现多个接口
type B interface {
	test1()
}
type C interface {
	test2()
}
type D interface {
	B
	C
	test3()
}

//A结构体实现上面的接口
func (a *A) test1() {

}
func (a *A) test2() {

}
func (a *A) test3() {

}

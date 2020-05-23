package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体	一个程序就是一个世界，有很多对象（变量）
	1.golang和传统的面向对象编程有区别，并不是纯粹的面向对象语言。
	2.Golang没有类，Go语言的结构体和其他编程语言的类有同等的地位，
	3.Golang面向对象编程非常简洁，去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
	4.Golang仍然有面向对象编程的继承、封装和多态的特性，只是实现的方式和其他OOP语言不一样，比如继承：Golang没有extends关键字，继承是通过匿名字段来实现。
	5.Golang面向对象很优雅，OOP本身就是语言类型系统的一部分，通过接口关联，耦合性低，也非常灵活。
结构体在传递时是值传递，想要改变原结构体使用地址
结构体的注意事项和使用细节
	1.结构体的所有字段在内存中是连续的，因此结构体的处理速度快
	2.如果结构体字段/属性是指针类型，那么指针类型本身的地址是连续的，它指向的地址不一定是连续的
	3.结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的属性（属性名，类型，个数都要完全相同）
	4.结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转。
	5.struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化。
		可以将结构体序列化为json字符串，使用encoding/json包中的Marshal(struct),该方法返回一个[]byte和一个错误err
*/
type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   int    `json:"id"`
}

func main() {
	monster := Monster{
		Name: "dog",
		Age:  23,
		Id:   111,
	}
	jsonStr, _ := json.Marshal(monster)
	fmt.Println(string(jsonStr))
}

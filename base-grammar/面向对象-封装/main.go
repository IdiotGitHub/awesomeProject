package main

import (
	"awesomeProject/base-grammar/面向对象-封装/modal"
	"fmt"
)

/*
面向对象编程思想--抽象
	我们在定义一个结构体的时候，实际上就是把一类事物的工有的属性（字段）和行为（方法）提取出来，形成一个物理模型，这种研究问题的方法称为抽象
面向对象编程三大特征
	Golang仍然有面向对象编程的继承、封装、多态的特性，只是是想的方式和其他OOP语言不一样。
封装
	encapsulation：就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只能通过被授权的操作（方法），才能对字段进行操作。
	封装的好处
		1.隐藏实现细节
		2.可以对数据进行验证，保证安全合理
	如何体现封装
		1.对结构体中的属性进行封装（首字母小写）
		2.通过方法，包实现封装
实现封装的步骤
	1.将结构体、字段（属性）的首字母小写（不能导出，其他包不能使用，类似private）
	2.给结构体所在包提供一个工厂模式的函数，首字母大写，类似一个构造函数
	3.提供一个首字母大写的Set方法，用于堆属性判断并赋值
	4.提供一个首字母大写的Get方法，用于获取属性的值

*/

func main() {
	person1 := modal.NewPerson("smith")
	fmt.Println(person1.GetAge())
}

type Account struct {
	AccountNo int
	Password  string
	Balance   float64
}

func (a *Account) Deposit(money float64, password string) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}
	if money < 0 {
		fmt.Println("存款金额错误")
		return
	}
	a.Balance += money
	fmt.Println("存款成功，存款", money)
}

func (a *Account) Withdraw(money float64, password string) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}
	if money < 0 || money > a.Balance {
		fmt.Println("取款金额错误")
		return
	}
	a.Balance -= money
	fmt.Println("取款成功，余额", a.Balance)
}
func (a *Account) Query(password string) {
	if password != a.Password {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("查询成功，余额", a.Balance)
}

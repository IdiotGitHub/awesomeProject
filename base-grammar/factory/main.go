package main

import "fmt"
import "awesomeProject/base-grammar/factory/modal"

/*
工厂模式
	Golang的结构体没有构造函数通常可以使用工厂模式来解决这个问题
当我们定义的结构体首字母是小写而且我们又想在其他的包中使用这个结构体，就可以使用工厂模式解决

*/

func main() {
	/*student := modal.Student{Name: "tom", Score: 99}
	fmt.Println(student)*/
	student := modal.NewStudent("tom", 99)
	fmt.Println(*student)
}

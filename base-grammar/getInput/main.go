package main

import "fmt"

func main() {
	//使用fmt.Scanln()获取用户输入,读取一行输入
	var str int64
	fmt.Println("请输入一个整数")
	fmt.Scanln(&str)
	fmt.Println("你输入的整数为：", str)
	//也可以使用格式化输入fmt.Scanf()
	//exg：fmt.Scanf("%s,%d,%f",&str,&int,&float)输入的时候严格按照引号内的格式进行输入
}

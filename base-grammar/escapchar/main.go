package main

import "fmt"

func main() {
	//\r的作用是一个回车，将\r后面的字符替换前面的相同长度的字符
	fmt.Println("test123456789\rhello bitch")
	fmt.Println("test\nhello bitch")
	fmt.Println("test\thello bitch")
}

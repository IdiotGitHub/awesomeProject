package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
go语言循环控制,for循环：
	for循环在声明变量时不允许使用关键字var来进行声明，也就是不能指定特定的数据类型，想要规定类型就需要将变量声明与初始化放到for循环外面
	1.循环条件是返回一个布尔值的表达式
	2.for循环的第二种使用方式
		for 循环条件 {
			//循环执行语句
		}
		将变量初始化和变量迭代写在别的地方
	3.for循环第三种写法,通常配合break使用
		for { //等价于for ;; {}
			//循环执行语句
		}
	4.for-range
		for index, value := range var {}
		如果在遍历字符串的时候字符串中存在中文字符，需要注意，需要先把字符串转为切片来进行遍历。[]rune(str)
		因为在使用range进行遍历的时候会按照字节进行
编程思想，将整个需求进行拆分，化整为零、各个击破
		先死后活，先把参数写死、再参数化
跳转控制语句-break
	1.break语句出现在多层嵌套的语句块中时，可以通过标签知名要终止的是哪一层语句块
	2.如果break后面不加标签名，那么break只终止最近的语句块
	3.使用标签的话，是表明break要终止的是哪一层循环，而不是跳转到标签处。
continue
	1.continue 接数本次循环进入下次循环
	2.continue支持标签，接数该标签标记的当前循环进入该循环的下次循环
goto
	无条件的跳转到某个代码块中
return
	使用在方法或者函数中，表示跳出所在的方法或函数；如果return在mian函数中，表示退出程序
*/

func main() {
	//forControl()
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	exercise6()
}

func forControl() {
	/*for i := 1; i < 10; i++ {
		fmt.Println("hello,world!")
	}*/
	//for-range遍历字符串
	var str1 string = "hello world"
	//传统的方式对字符串进行遍历
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	//使用for-range进行遍历
	for index, val := range str1 {
		fmt.Printf("index = %d, val = %c\n", index, val)
	}
	//当字符串中存在中文字符或者存在字节数大于一的字符的时候需要先将字符串转为切片
	var str2 string = "北极星"
	str3 := []rune(str2)
	for index, value := range str3 {
		fmt.Printf("index = %d, val = %c\n", index, value)

	}
}
func exercise1() {
	var count int
	var sum int
	for i := 1; i < 101; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("一共有%d个9的倍数，它们的和为%d\n", count, sum)
}

func exercise2() {
	for i := 0; i < 7; i++ {
		fmt.Printf("%d + %d = 6\n", i, 6-i)
	}
}
func exercise3() {
	var score float64
	var all = 0.0
	var class int32 = 2
	var student int32 = 3
	for i := 1; i <= int(class); i++ {
		var sum = 0.0
		score = 0
		for j := 1; j <= int(student); j++ {
			fmt.Printf("请输入第%d个班级,第%d个同学的成绩：\n", i, j)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("第%d个班级的平均分为%f\n", i, sum/float64(student))
		all += sum
	}
	fmt.Println("所有同学的平均成绩为", float64(all/2/3))
}
func exercise4() {
	//打印空心金字塔
	//层数
	var totalLevel int = 9
	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if i == totalLevel || j == 1 || j == i*2-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func exercise5() {
	//随机生成[0,n)的整数使用rand.Intn(n),但是如果不设置一个种子的话，他每次随机的数字都是相同的
	var count int
	for true {
		count++
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println(num)
		if num == 99 {
			break
		}
	}
	fmt.Println(count)
}

func exercise6() {
	//break练习
	rand.Seed(time.Now().UnixNano())
	for {
		num1 := rand.Intn(100) + 1
		num2 := rand.Intn(100) + 1
		if num1+num2 > 20 {
			fmt.Printf("这两个数分别为%d、%d\n", num1, num2)
			break
		}
	}
}

func exercise7() {

}
